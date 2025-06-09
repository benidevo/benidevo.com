package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/benidevo/website/internal/config"
)

// GitHubClient provides common GitHub API functionality
type GitHubClient struct {
	cfg        *config.GitHubConfig
	httpClient *http.Client
}

// HTTPClient returns the HTTP client for external use
func (c *GitHubClient) HTTPClient() *http.Client {
	return c.httpClient
}

// GitHubFile represents a file response from GitHub Contents API
type GitHubFile struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
	DownloadURL string `json:"download_url"`
}

// NewGitHubClient creates a new GitHub client
func NewGitHubClient(cfg *config.GitHubConfig) *GitHubClient {
	return &GitHubClient{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// FetchFileContent fetches and decodes file content from GitHub repository
func (c *GitHubClient) FetchFileContent(filePath string) (string, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/contents/%s",
		c.cfg.BaseURL, c.cfg.Owner, c.cfg.Repository, filePath)

	file, err := c.fetchGitHubFile(url)
	if err != nil {
		return "", err
	}

	return c.decodeFileContent(file)
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

// FetchRepositoryData fetches repository metadata from GitHub API
func (c *GitHubClient) FetchRepositoryData(owner, repo string) ([]byte, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", c.cfg.BaseURL, owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add authentication if token is provided
	if c.cfg.Token != "" {
		req.Header.Set("Authorization", "token "+c.cfg.Token)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "Portfolio-Website/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch repository data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API returned status %d: %s", resp.StatusCode, string(body))
	}

	return io.ReadAll(resp.Body)
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
