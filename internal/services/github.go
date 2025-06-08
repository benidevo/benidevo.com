package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/benidevo/website/internal/models"
)

// GitHubService handles GitHub API interactions
type GitHubService struct {
	baseURL    string
	username   string
	token      string
	httpClient *http.Client
	cache      map[string]cacheEntry
}

// cacheEntry represents a single entry in the cache, storing the cached data,
// the time it was cached, and its time-to-live duration.
type cacheEntry struct {
	data      interface{}
	timestamp time.Time
	ttl       time.Duration
}

// NewGitHubService creates a new GitHub service instance
func NewGitHubService(username, token string) *GitHubService {
	return &GitHubService{
		baseURL:  "https://api.github.com",
		username: username,
		token:    token,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		cache: make(map[string]cacheEntry),
	}
}

// GetFileContent fetches a file's content from a GitHub repository
func (g *GitHubService) GetFileContent(repo, filePath string) ([]byte, error) {
	cacheKey := fmt.Sprintf("file_%s_%s_%s", g.username, repo, filePath)

	// Check cache first
	if cached, exists := g.getFromCache(cacheKey); exists {
		if content, ok := cached.([]byte); ok {
			return content, nil
		}
	}

	url := fmt.Sprintf("%s/repos/%s/%s/contents/%s", g.baseURL, g.username, repo, filePath)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if g.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("token %s", g.token))
	}
	req.Header.Set("Accept", "application/vnd.github.v3.raw")

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file content: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("file not found or GitHub API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Cache the result for 10 minutes
	g.setCache(cacheKey, body, 10*time.Minute)

	return body, nil
}

// GetRepository fetches a specific repository
func (g *GitHubService) GetRepository(repoName string) (*models.GitHubRepository, error) {
	cacheKey := fmt.Sprintf("repo_%s_%s", g.username, repoName)

	// Check cache first
	if cached, exists := g.getFromCache(cacheKey); exists {
		if repo, ok := cached.(*models.GitHubRepository); ok {
			return repo, nil
		}
	}

	url := fmt.Sprintf("%s/repos/%s/%s", g.baseURL, g.username, repoName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if g.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("token %s", g.token))
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch repository: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var repo models.GitHubRepository
	if err := json.Unmarshal(body, &repo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Cache the result for 5 minutes
	g.setCache(cacheKey, &repo, 5*time.Minute)

	return &repo, nil
}

// GetSkillsData fetches skills.json from the data repository
func (g *GitHubService) GetSkillsData(dataRepo string) ([]models.SkillCategory, error) {
	content, err := g.GetFileContent(dataRepo, "skills.json")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch skills data: %w", err)
	}

	var skills []models.SkillCategory
	if err := json.Unmarshal(content, &skills); err != nil {
		return nil, fmt.Errorf("failed to unmarshal skills data: %w", err)
	}

	return skills, nil
}

// Helper methods for caching
func (g *GitHubService) getFromCache(key string) (interface{}, bool) {
	entry, exists := g.cache[key]
	if !exists {
		return nil, false
	}

	// Check if cache entry has expired
	if time.Since(entry.timestamp) > entry.ttl {
		delete(g.cache, key)
		return nil, false
	}

	return entry.data, true
}

func (g *GitHubService) setCache(key string, data interface{}, ttl time.Duration) {
	g.cache[key] = cacheEntry{
		data:      data,
		timestamp: time.Now(),
		ttl:       ttl,
	}
}

// ClearCache removes all cached entries
func (g *GitHubService) ClearCache() {
	g.cache = make(map[string]cacheEntry)
}
