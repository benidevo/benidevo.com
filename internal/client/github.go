package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/benidevo/website/internal/config"
)

// CacheEntry represents a cached file with expiration
type CacheEntry struct {
	Content   string
	ExpiresAt time.Time
}

// GitHubClient provides common GitHub API functionality
type GitHubClient struct {
	cfg        *config.GitHubConfig
	httpClient *http.Client
	cache      map[string]CacheEntry
	cacheMutex sync.RWMutex
	cacheTTL   time.Duration
}

// HTTPClient returns the HTTP client for external use
func (c *GitHubClient) HTTPClient() *http.Client {
	return c.httpClient
}

// GitHubFile represents a file response from GitHub Contents API
type GitHubFile struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

// NewGitHubClient creates a new GitHub client
func NewGitHubClient(cfg *config.GitHubConfig) *GitHubClient {
	return &GitHubClient{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		cache:    make(map[string]CacheEntry),
		cacheTTL: 15 * time.Minute, // Cache for 15 minutes
	}
}

// FetchFileContent fetches and decodes file content from GitHub repository with caching
func (c *GitHubClient) FetchFileContent(filePath string) (string, error) {
	c.cacheMutex.RLock()
	entry, exists := c.cache[filePath]
	c.cacheMutex.RUnlock()

	if exists && time.Now().Before(entry.ExpiresAt) {
		return entry.Content, nil
	}

	url := fmt.Sprintf("%s/repos/%s/%s/contents/%s",
		c.cfg.BaseURL, c.cfg.Owner, c.cfg.Repository, filePath)

	file, err := c.fetchGitHubFile(url)
	if err != nil {
		return "", err
	}

	content, err := c.decodeFileContent(file)
	if err != nil {
		return "", err
	}

	c.cacheMutex.Lock()
	c.cache[filePath] = CacheEntry{
		Content:   content,
		ExpiresAt: time.Now().Add(c.cacheTTL),
	}
	c.cacheMutex.Unlock()

	return content, nil
}

// fetchGitHubFile makes a request to GitHub API and returns the file data
func (c *GitHubClient) fetchGitHubFile(url string) (*GitHubFile, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if c.cfg.Token != "" {
		req.Header.Set("Authorization", "token "+c.cfg.Token)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "Portfolio-Website/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API returned status %d: %s", resp.StatusCode, string(body))
	}

	var file GitHubFile
	if err := json.NewDecoder(resp.Body).Decode(&file); err != nil {
		return nil, fmt.Errorf("failed to decode GitHub response: %w", err)
	}

	return &file, nil
}

// decodeFileContent decodes base64 content from GitHub API response
func (c *GitHubClient) decodeFileContent(file *GitHubFile) (string, error) {
	if file.Encoding != "base64" {
		return "", fmt.Errorf("unsupported encoding: %s", file.Encoding)
	}

	// Remove any whitespace that might be in the base64 content
	content := file.Content
	content = removeWhitespace(content)

	decoded, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 content: %w", err)
	}

	return string(decoded), nil
}

// FetchBinaryFileContent fetches binary file content from GitHub repository and returns base64 string
func (c *GitHubClient) FetchBinaryFileContent(filePath string) (string, error) {
	c.cacheMutex.RLock()
	entry, exists := c.cache[filePath]
	c.cacheMutex.RUnlock()

	if exists && time.Now().Before(entry.ExpiresAt) {
		return entry.Content, nil
	}

	url := fmt.Sprintf("%s/repos/%s/%s/contents/%s",
		c.cfg.BaseURL, c.cfg.Owner, c.cfg.Repository, filePath)

	file, err := c.fetchGitHubFile(url)
	if err != nil {
		return "", err
	}

	if file.Encoding != "base64" {
		return "", fmt.Errorf("unsupported encoding for binary file: %s", file.Encoding)
	}

	content := removeWhitespace(file.Content)

	c.cacheMutex.Lock()
	c.cache[filePath] = CacheEntry{
		Content:   content,
		ExpiresAt: time.Now().Add(c.cacheTTL),
	}
	c.cacheMutex.Unlock()

	return content, nil
}

// ClearExpiredCache removes expired entries from cache
func (c *GitHubClient) ClearExpiredCache() {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()

	now := time.Now()
	for key, entry := range c.cache {
		if now.After(entry.ExpiresAt) {
			delete(c.cache, key)
		}
	}
}

// removeWhitespace removes all whitespace characters from a string
func removeWhitespace(s string) string {
	var result []byte
	for _, b := range []byte(s) {
		if b != ' ' && b != '\t' && b != '\n' && b != '\r' {
			result = append(result, b)
		}
	}
	return string(result)
}
