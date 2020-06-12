package config

type WebServerSettings struct {
	Scheme   string `yaml:"scheme"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
