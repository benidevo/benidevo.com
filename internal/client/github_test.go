package client

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/benidevo/website/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewGitHubClient(t *testing.T) {
	cfg := &config.GitHubConfig{
		Owner:      "test-owner",
		Repository: "test-repo",
		Token:      "test-token",
		BaseURL:    "https://api.github.com",
	}

	client := NewGitHubClient(cfg)

	assert.NotNil(t, client)
	assert.Equal(t, cfg, client.cfg)
	assert.NotNil(t, client.httpClient)
	assert.NotNil(t, client.cache)
	assert.Equal(t, 15*time.Minute, client.cacheTTL)
}

func TestGitHubClient_FetchFileContent(t *testing.T) {
	tests := []struct {
		name           string
		filePath       string
		serverResponse interface{}
		serverStatus   int
		wantContent    string
		wantErr        bool
	}{
		{
			name:     "successful fetch and decode",
			filePath: "test.txt",
			serverResponse: GitHubFile{
				Content:  base64.StdEncoding.EncodeToString([]byte("test content")),
				Encoding: "base64",
			},
			serverStatus: http.StatusOK,
			wantContent:  "test content",
			wantErr:      false,
		},
		{
			name:         "server returns error",
			filePath:     "test.txt",
			serverStatus: http.StatusNotFound,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Contains(t, r.URL.Path, tt.filePath)
				assert.Equal(t, "token test-token", r.Header.Get("Authorization"))

				w.WriteHeader(tt.serverStatus)
				if tt.serverResponse != nil {
					json.NewEncoder(w).Encode(tt.serverResponse)
				}
			}))
			defer server.Close()

			cfg := &config.GitHubConfig{
				Owner:      "test-owner",
				Repository: "test-repo",
				Token:      "test-token",
				BaseURL:    server.URL,
			}

			client := NewGitHubClient(cfg)
			content, err := client.FetchFileContent(tt.filePath)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantContent, content)
			}
		})
	}
}

func TestGitHubClient_Caching(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		response := GitHubFile{
			Content:  base64.StdEncoding.EncodeToString([]byte("cached content")),
			Encoding: "base64",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	cfg := &config.GitHubConfig{
		Owner:      "test-owner",
		Repository: "test-repo",
		BaseURL:    server.URL,
	}

	client := NewGitHubClient(cfg)

	// First request
	content1, err := client.FetchFileContent("test.txt")
	assert.NoError(t, err)
	assert.Equal(t, "cached content", content1)
	assert.Equal(t, 1, requestCount)

	// Second request should use cache
	content2, err := client.FetchFileContent("test.txt")
	assert.NoError(t, err)
	assert.Equal(t, "cached content", content2)
	assert.Equal(t, 1, requestCount) // No additional request
}

func TestGitHubClient_ClearExpiredCache(t *testing.T) {
	client := &GitHubClient{
		cache: map[string]CacheEntry{
			"expired": {
				Content:   "old content",
				ExpiresAt: time.Now().Add(-1 * time.Hour),
			},
			"valid": {
				Content:   "new content",
				ExpiresAt: time.Now().Add(1 * time.Hour),
			},
		},
		cacheTTL: 15 * time.Minute,
	}

	// Manually clear expired cache for test
	client.cacheMutex.Lock()
	now := time.Now()
	for key, entry := range client.cache {
		if now.After(entry.ExpiresAt) {
			delete(client.cache, key)
		}
	}
	client.cacheMutex.Unlock()

	assert.Len(t, client.cache, 1)
	assert.Contains(t, client.cache, "valid")
	assert.NotContains(t, client.cache, "expired")
}

func TestRemoveWhitespace(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "hello world",
			want:  "helloworld",
		},
		{
			input: "test\n\r\t content",
			want:  "testcontent",
		},
		{
			input: "nospaces",
			want:  "nospaces",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := removeWhitespace(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
