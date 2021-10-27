package configs

import (
	"encoding/json"
	"github.com/jfrog/jfrog-cli-core/v2/utils/coreutils"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"io/ioutil"
	"path/filepath"
)

var (
	globalConfFileName   = "jfrog.cli.bump.conf"
	dependenciesFileName = "bump.dependencies"
)

type VcsConfig struct {
	Url   string
	Token string
}
type DepConfig struct {
	Path       string
	Dependency string
}
type ConfigService interface {
	SaveVcs(VcsConfig) error
	ReadVcs() (VcsConfig, error)
	AddDep(DepConfig) error
	ReadDeps() ([]DepConfig, error)
}

type defaultConfigService struct {
}

func NewVcsConfigService() defaultConfigService {
	return defaultConfigService{}
}

func (defaultConfigService) ReadVcs() (VcsConfig, error) {
	filePath, err := getConfFilePath()
	if err != nil {
		return VcsConfig{}, err
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return VcsConfig{}, err
	}
	res := &VcsConfig{}
	err = json.Unmarshal(content, res)
	if err != nil {
		return VcsConfig{}, err
	}
	return *res, nil
}

func (defaultConfigService) SaveVcs(vcsConfig VcsConfig) error {
	confFilePath, err := getConfFilePath()
	if err != nil {
		return err
	}
	content, err := json.Marshal(vcsConfig)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(confFilePath, content, 0600)
	if err != nil {
		return err
	}
	log.Info("Configuration saved to " + confFilePath)
	return err
}

func (d defaultConfigService) AddDep(dep DepConfig) error {
	read, err := d.ReadDeps()
	for _, oneDep := range read {
		if dep == oneDep {
			return nil
		}
	}
	read = append(read, dep)
	content, err := json.Marshal(read)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dependenciesFileName, content, 0600)
}

func (d defaultConfigService) ReadDeps() ([]DepConfig, error) {
	content, err := ioutil.ReadFile(dependenciesFileName)
	if err != nil {
		return nil, err
	}
	res := make([]DepConfig, 0)
	err = json.Unmarshal(content, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func getConfFilePath() (string, error) {
	homeDir, err := coreutils.GetJfrogHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, globalConfFileName), nil
}
