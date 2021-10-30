package model

type PullRequest struct {
	// Pull request title
	Title string
	// Pull request description
	Description string
	// "from", source branch for pull request
	Head string
	// "to", target branch for pull request
	Base string
}
