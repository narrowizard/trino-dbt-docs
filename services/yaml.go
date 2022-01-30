package services

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func WriteToYaml(data interface{}, filename string) error {
	d, err := yaml.Marshal(&data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, d, 0644)
}
