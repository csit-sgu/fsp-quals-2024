package config

type Databases struct {
	Clickhouse DatabaseConfig `yaml:"clickhouse"`
	Postgres   DatabaseConfig `yaml:"postgres"`
}

type DatabaseConfig struct {
	Username string `yaml:"username"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}
