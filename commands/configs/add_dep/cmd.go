package add_dep

import (
	"fmt"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	log2 "github.com/jfrog/jfrog-cli-core/v2/utils/log"
	"github.com/jfrog/jfrog-cli-plugin-template/commands/configs"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

func init() {
	log2.SetDefaultLogger()
}

func GetAddDepCommand() components.Command {
	return components.Command{
		Name:        "add-dep",
		Description: "Declare a new dependency",
		Aliases:     []string{"d"},
		Arguments:   getDepArguments(),
		Action: func(c *components.Context) error {
			service := addDepService{
				configService:   configs.NewVcsConfigService(),
				bitbucketClient: configs.NewBitbucketClient(),
			}
			return service.addDepCmd(c)
		},
	}
}

func getDepArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "url",
			Description: "The bitbucket url of the file where the dependency is used",
		},
		{
			Name:        "package name",
			Description: "The dependency package name",
		},
	}
}

type addDepService struct {
	configService   configs.ConfigService
	bitbucketClient configs.BitbucketClient
}

func (s addDepService) addDepCmd(c *components.Context) error {
	// validate params
	if len(c.Arguments) != 2 {
		return fmt.Errorf("2 arguments requires : url, dependency")
	}
	path := c.Arguments[0]
	if path == "" {
		return fmt.Errorf("path required")
	}
	dependency := c.Arguments[1]
	if dependency == "" {
		return fmt.Errorf("dependency required")
	}
	config := configs.DepConfig{
		Path:       path,
		Dependency: dependency,
	}
	vcs, err := s.configService.ReadVcs()
	if err != nil {
		return err
	}
	err = s.bitbucketClient.Exists(vcs.Url, vcs.Token, path)
	if err != nil {
		return err
	}
	err = s.configService.AddDep(config)
	if err != nil {
		return err
	}
	log.Info("Dependency added")
	return nil
}
