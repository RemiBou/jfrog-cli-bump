package commands

import (
	"fmt"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	fakeVcsProviderConnectionChecker = &fakeDefaultVcsProviderConnectionCheckerImpl{}
	fakeVcsConfigSaver               = &fakeVcsConfigSaverImpl{}
)

func init() {
	defaultVcsProviderConnectionChecker = fakeVcsProviderConnectionChecker
	defaultVcsProviderConfigSaver = fakeVcsConfigSaver
}

func TestVcsCmdRequire2Arguments(t *testing.T) {
	err := vcsCmd(&components.Context{Arguments: []string{}})
	require.Error(t, err)
}
func TestVcsCmdRequireUrl(t *testing.T) {
	err := vcsCmd(&components.Context{Arguments: []string{
		"",
		"token",
	}})
	require.Error(t, err)
}

func TestVcsCmdRequireToken(t *testing.T) {
	err := vcsCmd(&components.Context{Arguments: []string{
		"https://google.com",
		"",
	}})
	require.Error(t, err)
}

func TestVcsCmdCheckConnectionReturnErr(t *testing.T) {
	expected := fmt.Errorf("an error")
	fakeVcsProviderConnectionChecker.nextResult = expected
	err := vcsCmd(&components.Context{Arguments: []string{
		"https://google.com",
		"my-token",
	}})
	require.Equal(t, expected, err)
	require.Equal(t, "https://google.com", fakeVcsProviderConnectionChecker.lastUrl)
	require.Equal(t, "my-token", fakeVcsProviderConnectionChecker.lastToken)
}

func TestVcsCmdSavesInJfrogCliCfg(t *testing.T) {
	err := vcsCmd(&components.Context{Arguments: []string{
		"https://google.com",
		"my-token",
	}})
	require.NoError(t, err)
	require.Equal(t, "https://google.com", fakeVcsConfigSaver.lastUrl)
	require.Equal(t, "my-token", fakeVcsConfigSaver.lastToken)
}

type fakeDefaultVcsProviderConnectionCheckerImpl struct {
	nextResult error
	lastUrl    string
	lastToken  string
}

func (f *fakeDefaultVcsProviderConnectionCheckerImpl) check(url string, token string) error {
	err := f.nextResult
	f.lastUrl = url
	f.lastToken = token
	f.nextResult = nil
	return err
}

type fakeVcsConfigSaverImpl struct {
	nextResult error
	lastUrl    string
	lastToken  string
}

func (f *fakeVcsConfigSaverImpl) save(url string, token string) error {
	err := f.nextResult
	f.lastUrl = url
	f.lastToken = token
	f.nextResult = nil
	return err
}
