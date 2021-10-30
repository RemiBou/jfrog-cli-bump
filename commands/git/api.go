package git

import "github.com/jfrog/jfrog-cli-bump/commands/git/model"

type GitServer interface {
	CreatePullRequest(projectKey string, repoName string, pr model.PullRequest) error
}