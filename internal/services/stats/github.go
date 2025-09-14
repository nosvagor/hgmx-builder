package stats

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Github struct {
	HTTP *http.Client
}

type githubRepoResponse struct {
	Stargazers  int `json:"stargazers_count"`
	Forks       int `json:"forks_count"`
	OpenIssues  int `json:"open_issues_count"`
	Subscribers int `json:"subscribers_count"`
	Size        int `json:"size"`
}

func (g Github) client() *http.Client {
	if g.HTTP != nil {
		return g.HTTP
	}
	return &http.Client{Timeout: 5 * time.Second}
}

func (g Github) FetchRepoStats(ctx context.Context, owner, repo string) (RepoStat, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return RepoStat{}, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	res, err := g.client().Do(req)
	if err != nil {
		return RepoStat{}, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return RepoStat{}, fmt.Errorf("github api: %s", res.Status)
	}

	var gr githubRepoResponse
	if err := json.NewDecoder(res.Body).Decode(&gr); err != nil {
		return RepoStat{}, err
	}

	rs := RepoStat{
		Owner:       owner,
		Repo:        repo,
		Stars:       gr.Stargazers,
		Forks:       gr.Forks,
		OpenIssues:  gr.OpenIssues,
		Subscribers: gr.Subscribers,
		Size:        gr.Size,
		FetchedAt:   time.Now().UTC(),
	}
	return UpsertRepoStats(ctx, rs)
}

func (g Github) GetRepoStats(ctx context.Context, owner, repo string, maxAge time.Duration) (RepoStat, error) {
	rs, err := GetCachedRepoStats(ctx, owner, repo)
	if err == nil && !Stale(rs, maxAge) {
		return rs, nil
	}
	return g.FetchRepoStats(ctx, owner, repo)
}
