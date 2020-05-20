package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config - main structure for the yaml file
type Config struct {
	Domain struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
}

// Parse - uses for getting and parsing the config file in yaml format
func (c *Config) Parse(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(f, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
