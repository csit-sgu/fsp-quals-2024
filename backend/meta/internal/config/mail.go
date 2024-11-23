package config

type MailConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Subject  string `yaml:"subject"`
	Body     string `yaml:"body"`
}
