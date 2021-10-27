package configs

import (
	"github.com/jfrog/jfrog-cli-core/v2/utils/coreutils"
	log2 "github.com/jfrog/jfrog-cli-core/v2/utils/log"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func init() {
	log2.SetDefaultLogger()
}
func TestSaveVcsReadVcs(t *testing.T) {
	dir, err := ioutil.TempDir("", "bump-conf")
	defer os.RemoveAll(dir)
	require.NoError(t, err)
	_ = os.Setenv(coreutils.HomeDir, dir)
	defer func() {
		_ = os.Setenv(coreutils.HomeDir, "")
	}()

	expected := VcsConfig{
		Url:   "https://google.com",
		Token: "my-token",
	}
	saverImpl := defaultConfigService{}
	err = saverImpl.SaveVcs(expected)

	require.NoError(t, err)
	actual, err := saverImpl.ReadVcs()
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestSaveVcsReturnErrIfFileWriteErr(t *testing.T) {
	// use a non existent directory so write should fail
	dir := "/azzezeazea/azeeaze"
	_ = os.Setenv(coreutils.HomeDir, dir)
	defer func() {
		_ = os.Setenv(coreutils.HomeDir, "")
	}()

	expected := VcsConfig{}
	err := defaultConfigService{}.SaveVcs(expected)

	require.Error(t, err)
}

func TestReadVcsShouldReturnErrorIfReadFail(t *testing.T) {
	dir, err := ioutil.TempDir("", "bump-conf")
	defer os.RemoveAll(dir)
	require.NoError(t, err)
	_ = os.Setenv(coreutils.HomeDir, dir)
	defer func() {
		_ = os.Setenv(coreutils.HomeDir, "")
	}()

	_, err = defaultConfigService{}.ReadVcs()

	require.Error(t, err)
}

func TestAddDepCreatesLocalFile(t *testing.T) {
	getwd, _ := os.Getwd()
	dir, err := ioutil.TempDir("", "bump-conf")
	defer func() {
		os.RemoveAll(dir)
		_ = os.Chdir(getwd)
	}()
	require.NoError(t, err)
	_ = os.Chdir(dir)

	saver := defaultConfigService{}
	readDir, err := ioutil.ReadDir(".")
	require.NoError(t, err)
	require.Equal(t, 0, len(readDir))
	err = saver.AddDep(DepConfig{
		Path:       "/path/go.mod",
		Dependency: "dependency",
	})

	require.NoError(t, err)
	readDir, err = ioutil.ReadDir(".")
	require.NoError(t, err)
	require.Equal(t, 1, len(readDir))
}

func TestReadDeps(t *testing.T) {
	getwd, _ := os.Getwd()
	dir, err := ioutil.TempDir("", "bump-conf")
	defer func() {
		os.RemoveAll(dir)
		_ = os.Chdir(getwd)
	}()
	require.NoError(t, err)
	_ = os.Chdir(dir)

	saver := defaultConfigService{}
	err = saver.AddDep(DepConfig{
		Path:       "/path/go.mod",
		Dependency: "dependency",
	})
	err = saver.AddDep(DepConfig{
		Path:       "/path/go2.mod",
		Dependency: "dependency2",
	})
	require.NoError(t, err)
	actual, err := saver.ReadDeps()

	require.NoError(t, err)
	expected := []DepConfig{
		{
			Path:       "/path/go.mod",
			Dependency: "dependency",
		},
		{
			Path:       "/path/go2.mod",
			Dependency: "dependency2",
		},
	}
	require.ElementsMatch(t, expected, actual)
}

func TestAddDepIgnoreDuplicates(t *testing.T) {
	getwd, _ := os.Getwd()
	dir, err := ioutil.TempDir("", "bump-conf")
	defer func() {
		os.RemoveAll(dir)
		_ = os.Chdir(getwd)
	}()
	require.NoError(t, err)
	_ = os.Chdir(dir)

	saver := defaultConfigService{}
	err = saver.AddDep(DepConfig{
		Path:       "/path/go.mod",
		Dependency: "dependency",
	})
	err = saver.AddDep(DepConfig{
		Path:       "/path/go.mod",
		Dependency: "dependency",
	})
	require.NoError(t, err)
	actual, err := saver.ReadDeps()

	require.NoError(t, err)
	expected := []DepConfig{
		{
			Path:       "/path/go.mod",
			Dependency: "dependency",
		},
	}
	require.ElementsMatch(t, expected, actual)
}
