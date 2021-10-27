package set_vcs

import (
	"fmt"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-plugin-template/commands/configs"
)

type setVcsService struct {
	configService   configs.ConfigService
	bitbucketClient configs.BitbucketClient
}

func GetSetVcsCommand() components.Command {
	return components.Command{
		Name:        "set-vcs",
		Description: "Configure the VCS connection for the bump plugin.",
		Aliases:     []string{"v"},
		Arguments:   getVcsArguments(),
		Action: func(c *components.Context) error {
			service := setVcsService{
				configService:   configs.NewVcsConfigService(),
				bitbucketClient: configs.NewBitbucketClient(),
			}
			return service.vcsCmd(c)
		},
	}
}

func getVcsArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "url",
			Description: "The bitbucket server url.",
		},
		{
			Name:        "username",
			Description: "The bitbucket username.",
		},
		{
			Name:        "token",
			Description: "The bitbucket token.",
		},
	}
}

func (s setVcsService) vcsCmd(c *components.Context) error {
	// validate params
	if len(c.Arguments) != 2 {
		return fmt.Errorf("2 arguments requires : url, token")
	}
	url := c.Arguments[0]
	if url == "" {
		return fmt.Errorf("url required")
	}
	token := c.Arguments[1]
	if token == "" {
		return fmt.Errorf("token required")
	}
	err := s.bitbucketClient.TestConnection(url, token)
	if err != nil {
		return err
	}
	err = s.configService.SaveVcs(configs.VcsConfig{
		Url:   url,
		Token: token,
	})
	if err != nil {
		return err
	}
	return err
}
