package configs

import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddDepCmdRequire2Arguments(t *testing.T) {
	service := addDepService{}
	err := service.addDepCmd(&components.Context{Arguments: []string{}})
	require.Error(t, err)
}

func TestAddDepCmdFailIfEmptyUrl(t *testing.T) {
	service := addDepService{}
	err := service.addDepCmd(&components.Context{Arguments: []string{"", "jfrog.com/my-dependency"}})
	require.Error(t, err)
}

func TestAddDepCmdFailIfEmptyDependency(t *testing.T) {
	service := addDepService{}
	err := service.addDepCmd(&components.Context{Arguments: []string{"https://google.com", ""}})
	require.Error(t, err)
}
