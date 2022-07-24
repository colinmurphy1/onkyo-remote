package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Parse configuration file
func (c *Config) parseConfig(configFile string) error {
	// Read yaml file
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	// Load yaml config into Config struct
	if err := yaml.UnmarshalStrict(content, c); err != nil {
		return err
	}

	return nil
}
