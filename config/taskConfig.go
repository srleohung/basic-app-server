package config

type ExampleFetchingTaskConfig struct {
	Period int `yaml:"period"`
}

type ExampleReportingTaskConfig struct {
	Period   int `yaml:"period"`
	PageSize int `yaml:"page_size"`
}

type TasksConfig struct {
	ExampleFetchingTask  *ExampleFetchingTaskConfig  `yaml:"example_fetching_task"`
	ExampleReportingTask *ExampleReportingTaskConfig `yaml:"example_reporting_task"`
}
