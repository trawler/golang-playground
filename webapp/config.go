package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Page comment
type Page struct {
	Title string
	Body  []byte
	Path  string `json:",inline" yaml:"path,omitempty"`
	User  string `json:",inline" yaml:"user,omitempty"`
}

// SiteConfig comment
type SiteConfig struct {
	Site []Page `json:",inline" yaml:"configuration,omitempty"`
}

func parseYamlFile(path string) (*SiteConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", path)
	}
	return parseYaml(data)
}

func parseYaml(data []byte) (*SiteConfig, error) {
	config := &SiteConfig{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unable to parse yaml data file")
	}
	return config, nil
}
