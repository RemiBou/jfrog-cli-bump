package add_dep

import (
	"encoding/json"
	"io/ioutil"
)

var filename = "bump.dependencies"

type depSaver interface {
	add(depConfig) error
	read() ([]depConfig, error)
}
type defaultDepSaver struct {
}

func (d defaultDepSaver) add(dep depConfig) error {
	read, err := d.read()
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
	return ioutil.WriteFile(filename, content, 0600)
}

func (d defaultDepSaver) read() ([]depConfig, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	res := make([]depConfig, 0)
	err = json.Unmarshal(content, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
