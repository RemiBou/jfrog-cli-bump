package configs

import (
	"encoding/json"
	"github.com/jfrog/jfrog-cli-core/v2/utils/coreutils"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"io/ioutil"
	"path/filepath"
)

var (
	confFileName = "jfrog.cli.bump.conf"
)

type vcsConfigSaver interface {
	save(vcsConfig) error
	read() (vcsConfig, error)
}

type defaultVcsConfigSaver struct {
}

func (i defaultVcsConfigSaver) read() (vcsConfig, error) {
	filePath, err := getConfFilePath()
	if err != nil {
		return vcsConfig{}, err
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return vcsConfig{}, err
	}
	res := &vcsConfig{}
	err = json.Unmarshal(content, res)
	if err != nil {
		return vcsConfig{}, err
	}
	log.Info("Configuration saved to " + filePath)
	return *res, nil
}

func (defaultVcsConfigSaver) save(vcsConfig vcsConfig) error {
	confFilePath, err := getConfFilePath()
	if err != nil {
		return err
	}
	content, err := json.Marshal(vcsConfig)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(confFilePath, content, 0600)
}

func getConfFilePath() (string, error) {
	homeDir, err := coreutils.GetJfrogHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, confFileName), nil
}
