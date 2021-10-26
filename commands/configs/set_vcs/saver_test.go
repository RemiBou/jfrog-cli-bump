package set_vcs

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
func TestSaveRead(t *testing.T) {
	dir, err := ioutil.TempDir("", "bump-conf")
	defer os.RemoveAll(dir)
	require.NoError(t, err)
	_ = os.Setenv(coreutils.HomeDir, dir)
	defer func() {
		_ = os.Setenv(coreutils.HomeDir, "")
	}()

	expected := vcsConfig{
		Url:   "https://google.com",
		Token: "my-token",
	}
	saverImpl := defaultVcsConfigSaver{}
	err = saverImpl.save(expected)

	require.NoError(t, err)
	actual, err := saverImpl.read()
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestSaveReturnErrIfFileWriteErr(t *testing.T) {
	// use a non existent directory so write should fail
	dir := "/azzezeazea/azeeaze"
	_ = os.Setenv(coreutils.HomeDir, dir)
	defer func() {
		_ = os.Setenv(coreutils.HomeDir, "")
	}()

	expected := vcsConfig{}
	err := defaultVcsConfigSaver{}.save(expected)

	require.Error(t, err)
}

func TestReadShouldReturnErrorIfReadFail(t *testing.T) {
	dir, err := ioutil.TempDir("", "bump-conf")
	defer os.RemoveAll(dir)
	require.NoError(t, err)
	_ = os.Setenv(coreutils.HomeDir, dir)
	defer func() {
		_ = os.Setenv(coreutils.HomeDir, "")
	}()

	_, err = defaultVcsConfigSaver{}.read()

	require.Error(t, err)
}
