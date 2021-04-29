package shared

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Website struct {
	Name  string `yaml:"name"  json:"name"`
	Url   string `yaml:"url"   json:"url"`
	Regex string `yaml:"regex" json:"regex"`
}

type Checker struct {
	Interval int       `yaml:"interval"`
	Timeout  int       `yaml:"timeout"`
	Websites []Website `yaml:"websites"`
}

type Kafka struct {
	Address   string `yaml:"address"`
	Topic     string `yaml:"topic"`
	Partition int    `yaml:"partition"`
}

type Config struct {
	Checker Checker `yaml:"checker"`
	Kafka   Kafka   `yaml:"kafka"`
}

func ConfigFromEnv(name string) (*Config, error) {
	filename := os.Getenv(name)
	if filename == "" {
		return nil, fmt.Errorf("env variable %s is undefined", name)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read file %s: %w", filename, err)
	}

	config := new(Config)
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("unable to parse yaml: %w", err)
	}

	// TODO: validate the config!

	return config, nil
}
