package add_dep

import (
	"fmt"
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
	err := service.addDepCmd(&components.Context{Arguments: []string{"https://google.com", ""}})
	require.Error(t, err)
}

func TestAddDepSave(t *testing.T) {
	service := createServiceWithFakeDep()
	err := service.addDepCmd(&components.Context{Arguments: []string{"https://google.com", "jfrog.com/my-dependency"}})
	require.NoError(t, err)
	require.Equal(t, depConfig{
		url:        "https://google.com",
		dependency: "jfrog.com/my-dependency",
	}, service.saver.(*fakeSaver).lastParam)
}

func TestAddDepReturnErrIfSaveErr(t *testing.T) {
	service := createServiceWithFakeDep()
	expected := fmt.Errorf("an error")
	service.saver.(*fakeSaver).nextErr = expected
	err := service.addDepCmd(&components.Context{Arguments: []string{"https://google.com", "jfrog.com/my-dependency"}})
	require.Equal(t, expected, err)
}

func createServiceWithFakeDep() addDepService {
	return addDepService{
		saver: &fakeSaver{},
	}
}

type fakeSaver struct {
	nextErr   error
	lastParam depConfig
	nextRead  []depConfig
}

func (f *fakeSaver) add(config depConfig) error {
	f.lastParam = config
	err := f.nextErr
	f.nextErr = nil
	return err
}

func (f *fakeSaver) read() ([]depConfig, error) {
	panic("implement me")
}
