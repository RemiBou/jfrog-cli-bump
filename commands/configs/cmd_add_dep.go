package configs

import (
	"fmt"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
)

func GetAddDepCommand() components.Command {
	return components.Command{
		Name:        "add-dep",
		Description: "Declare a new dependency",
		Aliases:     []string{"d"},
		Arguments:   getDepArguments(),
		Action: func(c *components.Context) error {
			service := addDepService{}
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
}

func (s addDepService) addDepCmd(c *components.Context) error {
	// validate params
	if len(c.Arguments) != 2 {
		return fmt.Errorf("2 arguments requires : url, dependency")
	}
	url := c.Arguments[0]
	if url == "" {
		return fmt.Errorf("url required")
	}
	dependency := c.Arguments[1]
	if dependency == "" {
		return fmt.Errorf("dependency required")
	}
	return nil
}
