package set_vcs

import (
	"fmt"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-plugin-template/commands/configs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVcsCmdRequire2Arguments(t *testing.T) {
	service := createServiceWithFakeDeps()
	err := service.vcsCmd(&components.Context{Arguments: []string{}})
	require.Error(t, err)
	require.Equal(t, configs.VcsConfig{}, service.configService.(*configs.FakeConfigService).LastSaveVcsParam)
}

func createServiceWithFakeDeps() setVcsService {
	return setVcsService{
		configService:   &configs.FakeConfigService{},
		bitbucketClient: &configs.FakeClient{},
	}
}
func TestVcsCmdRequireUrl(t *testing.T) {
	service := createServiceWithFakeDeps()
	err := service.vcsCmd(&components.Context{Arguments: []string{
		"",
		"token",
	}})
	require.Error(t, err)
}

func TestVcsCmdRequireToken(t *testing.T) {
	service := createServiceWithFakeDeps()
	err := service.vcsCmd(&components.Context{Arguments: []string{
		"https://google.com",
		"",
	}})
	require.Error(t, err)
}

func TestVcsCmdCheckConnectionReturnErr(t *testing.T) {
	service := createServiceWithFakeDeps()
	expected := fmt.Errorf("an error")
	client := service.bitbucketClient.(*configs.FakeClient)
	client.NextErr = expected
	err := service.vcsCmd(&components.Context{Arguments: []string{
		"https://google.com",
		"my-token",
	}})
	require.Equal(t, expected, err)
	require.Equal(t, "https://google.com", client.LastParamTest.Url)
	require.Equal(t, "my-token", client.LastParamTest.Token)
}

func TestVcsCmdSavesInJfrogCliCfg(t *testing.T) {
	service := createServiceWithFakeDeps()
	err := service.vcsCmd(&components.Context{Arguments: []string{
		"https://google.com",
		"my-token",
	}})
	configSaver := service.configService.(*configs.FakeConfigService)
	require.NoError(t, err)
	require.Equal(t, "https://google.com", configSaver.LastSaveVcsParam.Url)
	require.Equal(t, "my-token", configSaver.LastSaveVcsParam.Token)
}
