package configs

import (
	"fmt"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
)

type vcsConfig struct {
	Url   string
	Token string
}

type setVcsService struct {
	saver   vcsConfigSaver
	checker vscConfigChecker
}

func GetSetVcsCommand() components.Command {
	return components.Command{
		Name:        "set-vcs",
		Description: "Configure the VCS connection for the bump plugin.",
		Aliases:     []string{"v"},
		Arguments:   getVcsArguments(),
		Action: func(c *components.Context) error {
			service := setVcsService{
				saver:   defaultVcsConfigSaver{},
				checker: defaultVcsConfigChecker{},
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
	err := s.checker.check(vcsConfig{
		Url:   url,
		Token: token,
	})
	if err != nil {
		return err
	}
	err = s.saver.save(vcsConfig{
		Url:   url,
		Token: token,
	})
	if err != nil {
		return err
	}
	return err
}
