package configs

import (
	"fmt"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVcsCmdRequire2Arguments(t *testing.T) {
	service := createServiceWithFakeDeps()
	err := service.vcsCmd(&components.Context{Arguments: []string{}})
	require.Error(t, err)
	require.Equal(t, vcsConfig{}, service.saver.(*fakeVcsConfigSaver).lastParam)
}

func createServiceWithFakeDeps() configService {
	return configService{
		saver:   &fakeVcsConfigSaver{},
		checker: &fakeVcsConfigChecker{},
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
	configChecker := service.checker.(*fakeVcsConfigChecker)
	configChecker.nextResult = expected
	err := service.vcsCmd(&components.Context{Arguments: []string{
		"https://google.com",
		"my-token",
	}})
	require.Equal(t, expected, err)
	require.Equal(t, "https://google.com", configChecker.lastUrl)
	require.Equal(t, "my-token", configChecker.lastToken)
}

func TestVcsCmdSavesInJfrogCliCfg(t *testing.T) {
	service := createServiceWithFakeDeps()
	err := service.vcsCmd(&components.Context{Arguments: []string{
		"https://google.com",
		"my-token",
	}})
	configSaver := service.saver.(*fakeVcsConfigSaver)
	require.NoError(t, err)
	require.Equal(t, "https://google.com", configSaver.lastParam.Url)
	require.Equal(t, "my-token", configSaver.lastParam.Token)
}

type fakeVcsConfigChecker struct {
	nextResult error
	lastUrl    string
	lastToken  string
}

func (f *fakeVcsConfigChecker) check(url string, token string) error {
	err := f.nextResult
	f.lastUrl = url
	f.lastToken = token
	f.nextResult = nil
	return err
}

type fakeVcsConfigSaver struct {
	nextErr   error
	nextRead  vcsConfig
	lastParam vcsConfig
}

func (f *fakeVcsConfigSaver) save(config vcsConfig) error {
	err := f.nextErr
	f.lastParam = config
	f.nextErr = nil
	return err
}

func (f *fakeVcsConfigSaver) read() (vcsConfig, error) {
	read := f.nextRead
	f.nextRead = vcsConfig{}
	f.nextErr = nil
	return read, f.nextErr
}
