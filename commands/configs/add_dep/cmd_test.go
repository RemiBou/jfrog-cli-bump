package add_dep

import (
	"fmt"
	"github.com/jfrog/jfrog-cli-bump/commands/configs"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddDepCmdRequire2Arguments(t *testing.T) {
	service := createServiceWithFakeDep()
	err := service.addDepCmd(&components.Context{Arguments: []string{}})
	require.Error(t, err)
}

func TestAddDepCmdFailIfEmptyUrl(t *testing.T) {
	service := createServiceWithFakeDep()
	err := service.addDepCmd(&components.Context{Arguments: []string{"", "jfrog.com/my-dependency"}})
	require.Error(t, err)
}

func TestAddDepCmdFailIfEmptyDependency(t *testing.T) {
	service := createServiceWithFakeDep()
	err := service.addDepCmd(&components.Context{Arguments: []string{"/path/go.mod", ""}})
	require.Error(t, err)
}

func TestAddDepSave(t *testing.T) {
	service := createServiceWithFakeDep()
	err := service.addDepCmd(&components.Context{Arguments: []string{"/path/go.mod", "jfrog.com/my-dependency"}})
	require.NoError(t, err)
	require.Equal(t, configs.DepConfig{
		Path:       "/path/go.mod",
		Dependency: "jfrog.com/my-dependency",
	}, service.configService.(*configs.FakeConfigService).LastAddDepParam)
}

func TestAddDepFailsIfCheckFails(t *testing.T) {
	service := createServiceWithFakeDep()
	expected := fmt.Errorf("an error")
	service.bitbucketClient.(*configs.FakeClient).NextErr = expected
	service.configService.(*configs.FakeConfigService).NextReadVcs = configs.VcsConfig{
		Url:   "https://google.com",
		Token: "token",
	}
	err := service.addDepCmd(&components.Context{Arguments: []string{"/path/go.mod", "jfrog.com/my-dependency"}})
	require.Equal(t, expected, err)
	require.Equal(t, configs.ExistsParams{
		Url:   "https://google.com",
		Token: "token",
		Path:  "/path/go.mod",
	}, service.bitbucketClient.(*configs.FakeClient).LastParamExists)
	require.Equal(t, configs.DepConfig{}, service.configService.(*configs.FakeConfigService).LastAddDepParam)
}

func createServiceWithFakeDep() addDepService {
	return addDepService{
		configService:   &configs.FakeConfigService{},
		bitbucketClient: &configs.FakeClient{},
	}
}
