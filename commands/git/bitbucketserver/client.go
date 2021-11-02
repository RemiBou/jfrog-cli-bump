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

func NewClient(ctx context.Context) git.GitServer {
	bbClient := bitbucketv1.NewAPIClient(ctx, &bitbucketv1.Configuration{
		BasePath:      "https://git.jfrog.info/rest",
		Host:          "", //TODO
		Scheme:        "", //TODO
		DefaultHeader: map[string]string{"content-type": "Application/json"},
		UserAgent:     "JFrog CLI bump",
		HTTPClient:    http.DefaultClient, // TODO

	})
	return &Client{bbClient: bbClient}
}

func (c *Client) GetFile(projectKey, repoName, path string) (string, error) {
	response, err := c.bbClient.DefaultApi.GetContent_11(projectKey, repoName, path, map[string]interface{}{})
	if err != nil {
		fmt.Printf("failed to fetch file [%v] from project [%v] repo [%v] with the error: %+v", path, projectKey, repoName, err)
		return "", err
	}
	fmt.Printf("fetched file [%v] from project [%v] repo [%v] successfully", path, projectKey, repoName)
	fmt.Printf("file content [%v]", response.Payload) //TODO delete
	return string(response.Payload), nil
}

func (c *Client) CreateBranch(projectKey string, repoName string, pr model.PullRequest) error {
	//c.bbClient.DefaultApi.CreateBranch() // TODO
	panic("implement me")
}

func (c *Client) PutFile(projectKey string, repoName string, pr model.PullRequest) error {
	//c.bbClient.DefaultApi.EditFile() // TODO
	panic("implement me")
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
