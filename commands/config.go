package commands

import (
	"fmt"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-core/v2/utils/coreutils"
	"path/filepath"
)

var (
	defaultVcsProviderConnectionChecker vcsProviderConnectionChecker = defaultVcsProviderConnectionCheckerImpl{}
	defaultVcsProviderConfigSaver       vcsConfigSaver               = defaultVcsProviderConfigSaverImpl{}
)

type vcsProviderConnectionChecker interface {
	check(url string, token string) error
}

type defaultVcsProviderConnectionCheckerImpl struct {
}

func (defaultVcsProviderConnectionCheckerImpl) check(url string, token string) error {
	return nil
}

type vcsConfigSaver interface {
	save(url string, token string) error
}

type defaultVcsProviderConfigSaverImpl struct {
}

func (defaultVcsProviderConfigSaverImpl) save(url string, token string) error {
	confPath, err := coreutils.GetJfrogHomeDir()
	if err != nil {
		return err
	}
	confPath = filepath.Join(confPath, "jfrog.cli.bump.conf")
	return nil
}

func GetConfigCommand() components.Command {
	return components.Command{
		Name:        "vcs",
		Description: "Configure the VCS connection for the bump plugin.",
		Aliases:     []string{"v"},
		Arguments:   getVcsArguments(),
		Action: func(c *components.Context) error {
			return vcsCmd(c)
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

func vcsCmd(c *components.Context) error {
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
	err := defaultVcsProviderConnectionChecker.check(url, token)
	if err != nil {
		return err
	}
	return defaultVcsProviderConfigSaver.save(url, token)
}
