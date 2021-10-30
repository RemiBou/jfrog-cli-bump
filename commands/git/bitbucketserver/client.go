package bitbucketserver

import (
	"context"
	"fmt"
	bitbucketv1 "github.com/gfleury/go-bitbucket-v1"
	"github.com/jfrog/jfrog-cli-bump/commands/git"
	"github.com/jfrog/jfrog-cli-bump/commands/git/model"
	"net/http"
)

type Client struct {
	bbClient *bitbucketv1.APIClient
}

func (c *Client) CreatePullRequest(projectKey string, repoName string, pr model.PullRequest) error {
	response, err := c.bbClient.DefaultApi.CreatePullRequest(projectKey, repoName, bitbucketv1.PullRequest{
		Title:       pr.Title,
		Description: pr.Description,
		State:       "OPEN",
		Open:        true,
		FromRef: bitbucketv1.PullRequestRef{
			ID:           "",
			DisplayID:    "",
			LatestCommit: "",
			Repository: bitbucketv1.Repository{
				Slug:          "",
				ID:            0,
				Name:          "",
				ScmID:         "",
				State:         "",
				StatusMessage: "",
				Forkable:      false,
				Project:       nil,
				Public:        false,
				Links:         nil,
				Owner:         nil,
				Origin:        nil,
			},
		},
		ToRef: bitbucketv1.PullRequestRef{
			ID:           "",
			DisplayID:    "",
			LatestCommit: "",
			Repository: bitbucketv1.Repository{
				Slug:          "",
				ID:            0,
				Name:          "",
				ScmID:         "",
				State:         "",
				StatusMessage: "",
				Forkable:      false,
				Project:       nil,
				Public:        false,
				Links:         nil,
				Owner:         nil,
				Origin:        nil,
			},
		},
		Author: nil,
	})
	if err != nil {
		fmt.Printf("failed to create a pull request with the error: %+v", response)
		return err
	}
	fmt.Printf("created pull request successfuly: %+v", response)
	return nil
}

func NewClient(ctx context.Context) git.GitServer {
	bbClient := bitbucketv1.NewAPIClient(ctx, &bitbucketv1.Configuration{
		BasePath:      "",
		Host:          "git.jfrog.info", //TODO
		Scheme:        "https",          //TODO
		DefaultHeader: nil,
		UserAgent:     "JFrog CLI bump",
		HTTPClient:    http.DefaultClient, // TODO
	})
	return &Client{bbClient:bbClient}
}
