package plugins

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type TableType string

const (
	TableTypeSource     TableType = "source"
	TableTypeModel      TableType = "model"
	TableTypeSeed       TableType = "seed"
	TableTypeDeprecated TableType = "deprecated"
)

type TableTypeMapping map[string]TableType

type TableTypeSchema struct {
	Version    int      `yaml:"version"`
	Models     []string `yaml:"models"`
	Sources    []string `yaml:"sources"`
	Seeds      []string `yaml:"seeds"`
	Deprecated []string `yaml:"deprecated"`
}

func ReadTableTypeMapping(filename string) (TableTypeMapping, error) {
	var content, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var schema TableTypeSchema
	err = yaml.Unmarshal(content, &schema)
	if err != nil {
		return nil, err
	}
	var dest = make(TableTypeMapping)
	for _, v := range schema.Models {
		dest[v] = TableTypeModel
	}
	for _, v := range schema.Sources {
		dest[v] = TableTypeSource
	}
	for _, v := range schema.Deprecated {
		dest[v] = TableTypeDeprecated
	}
	for _, v := range schema.Seeds {
		dest[v] = TableTypeSeed
	}
	return dest, nil
}
