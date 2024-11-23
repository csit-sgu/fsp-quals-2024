package config

type Databases struct {
	Clickhouse DatabaseConfig   `yaml:"clickhouse"`
	Postgres   DatabaseConfig   `yaml:"postgres"`
	OpenSearch OpenSearchConfig `yaml:"opensearch"`
}

type DatabaseConfig struct {
	Username string `yaml:"username"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

type OpenSearchConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Index    string `yaml:"index"`
}
