package add_dep

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestSaveCreatesLocalFile(t *testing.T) {
	getwd, _ := os.Getwd()
	dir, err := ioutil.TempDir("", "bump-conf")
	defer func() {
		os.RemoveAll(dir)
		_ = os.Chdir(getwd)
	}()
	require.NoError(t, err)
	_ = os.Chdir(dir)

	saver := defaultDepSaver{}
	readDir, err := ioutil.ReadDir(".")
	require.NoError(t, err)
	require.Equal(t, 0, len(readDir))
	err = saver.add(depConfig{
		Url:        "https://google.com",
		Dependency: "dependency",
	})

	require.NoError(t, err)
	readDir, err = ioutil.ReadDir(".")
	require.NoError(t, err)
	require.Equal(t, 1, len(readDir))
}

func TestRead(t *testing.T) {
	getwd, _ := os.Getwd()
	dir, err := ioutil.TempDir("", "bump-conf")
	defer func() {
		os.RemoveAll(dir)
		_ = os.Chdir(getwd)
	}()
	require.NoError(t, err)
	_ = os.Chdir(dir)

	saver := defaultDepSaver{}
	err = saver.add(depConfig{
		Url:        "https://google.com",
		Dependency: "dependency",
	})
	err = saver.add(depConfig{
		Url:        "https://google2.com",
		Dependency: "dependency2",
	})
	require.NoError(t, err)
	actual, err := saver.read()

	require.NoError(t, err)
	expected := []depConfig{
		{
			Url:        "https://google.com",
			Dependency: "dependency",
		},
		{
			Url:        "https://google2.com",
			Dependency: "dependency2",
		},
	}
	require.ElementsMatch(t, expected, actual)
}

func TestSaveIgnoreDuplicates(t *testing.T) {
	getwd, _ := os.Getwd()
	dir, err := ioutil.TempDir("", "bump-conf")
	defer func() {
		os.RemoveAll(dir)
		_ = os.Chdir(getwd)
	}()
	require.NoError(t, err)
	_ = os.Chdir(dir)

	saver := defaultDepSaver{}
	err = saver.add(depConfig{
		Url:        "https://google.com",
		Dependency: "dependency",
	})
	err = saver.add(depConfig{
		Url:        "https://google.com",
		Dependency: "dependency",
	})
	require.NoError(t, err)
	actual, err := saver.read()

	require.NoError(t, err)
	expected := []depConfig{
		{
			Url:        "https://google.com",
			Dependency: "dependency",
		},
	}
	require.ElementsMatch(t, expected, actual)
}
