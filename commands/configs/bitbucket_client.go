package configs

import (
	"github.com/jfrog/froggit-go/vcsclient"
	"github.com/jfrog/froggit-go/vcsutils"
)

type BitbucketClient interface {
	TestConnection(url string, token string) error
	Exists(url string, token string, path string) error
}

type defaultBitbucketClient struct {
}

func NewBitbucketClient() defaultBitbucketClient {
	return defaultBitbucketClient{}
}

func (defaultBitbucketClient) TestConnection(url string, token string) error {
	// The VCS provider. Cannot be changed.
	client, err := vcsclient.NewClientBuilder(vcsutils.BitbucketServer).
		ApiEndpoint(url).
		Token(token).
		Build()
	if err != nil {
		return err
	}
	return client.TestConnection()
}

func (defaultBitbucketClient) Exists(url string, token string, path string) error {
	return nil
}
