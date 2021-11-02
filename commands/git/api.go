package git

import "github.com/jfrog/jfrog-cli-bump/commands/git/model"

type GitServer interface {
	GetFile(projectKey, repoName, path string) (string, error)
	CreateBranch(projectKey string, repoName string, pr model.PullRequest) error //TODO
	PutFile(projectKey string, repoName string, pr model.PullRequest) error      //TODO
	CreatePullRequest(projectKey string, repoName string, pr model.PullRequest) error
}
