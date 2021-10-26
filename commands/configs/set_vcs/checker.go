package set_vcs

import (
	"github.com/jfrog/froggit-go/vcsclient"
	"github.com/jfrog/froggit-go/vcsutils"
)

type vscConfigChecker interface {
	check(config vcsConfig) error
}

type defaultVcsConfigChecker struct {
}

func (defaultVcsConfigChecker) check(config vcsConfig) error {
	// The VCS provider. Cannot be changed.
	client, err := vcsclient.NewClientBuilder(vcsutils.BitbucketServer).
		ApiEndpoint(config.Url).
		Token(config.Token).
		Build()
	if err != nil {
		return err
	}
	return client.TestConnection()
}
