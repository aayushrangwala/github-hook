package github

import (
	"context"
	"log"

	githttp "github.com/google/go-github/v32/github"
	gitql "github.com/shurcooL/githubql"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type APIClient struct {
	gqlc   *gitql.Client
	client *githttp.Client

	token string

	ctx context.Context

	Logger *logrus.Logger
}

func NewAPIClient(token string) *APIClient {
	tc := &oauth2.Token{
		AccessToken: token,
	}

	httpCl := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(tc))
	if httpCl == nil {
		log.Fatalf("Nil oAuth client: %+v", httpCl)
	}

	return &APIClient{
		client: githttp.NewClient(httpCl),
		gqlc:   gitql.NewClient(httpCl),
		ctx:    context.Background(),
		token:  token,
		Logger: logrus.New(),
	}
}

func (c *APIClient) ListRepos() ([]*githttp.Repository, error) {
	opt := &githttp.RepositoryListOptions{
		Affiliation: "owner",
		ListOptions: githttp.ListOptions{PerPage: 5},
	}

	// get all pages of results
	var allRepos []*githttp.Repository
	for {
		repos, resp, err := c.client.Repositories.List(c.ctx, "", opt)
		if err != nil {
			return allRepos, err
		}

		allRepos = append(allRepos, repos...)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allRepos, nil
}
