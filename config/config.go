package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const DAFAULT_CONFIG_FILENAME string = "config.yaml"

type Config struct {
	LocalResourcePath string            `yaml:"local_resource_path"`
	DatastoreSettings DatastoreSettings `yaml:"datastore"`
	WebServerSettings WebServerSettings `yaml:"web_server"`
	Tasks             TasksConfig       `yaml:"tasks"`
}

var DEFAULT_CONFIG Config = Config{
	LocalResourcePath: "./",
	DatastoreSettings: DatastoreSettings{
		Type:     "",
		Host:     "localhost",
		Username: "root",
		Password: "root",
	},
	WebServerSettings: WebServerSettings{
		Scheme:   "http",
		Host:     "localhost",
		Username: "root",
		Password: "root",
	},
	Tasks: TasksConfig{
		ExampleFetchingTask:  &ExampleFetchingTaskConfig{Period: 60},
		ExampleReportingTask: &ExampleReportingTaskConfig{Period: 60, PageSize: 1},
	},
}

func LoadConfigFromYaml(filename string) *Config {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error("Cannot read config file '", filename, "'")
		return nil
	}
	var config = &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		logrus.Error("Cannot parse config file '", filename, "'")
		return nil
	}
	return config
}

func (config *Config) SaveConfigToYamlFile(filename string) error {
	if yamlBytes, err := yaml.Marshal(config); err != nil {
		return err
	} else {
		return ioutil.WriteFile(filename, yamlBytes, 0644)
	}
}
