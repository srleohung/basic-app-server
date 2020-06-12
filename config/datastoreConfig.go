package config

type DatastoreSettings struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
