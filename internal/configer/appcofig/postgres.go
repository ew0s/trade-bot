package appcofig

type PostgresConfiguration struct {
	Host        string              `yaml:"host"`
	Port        string              `yaml:"port"`
	Credentials PostgresCredentials `yaml:"credentials"`
	DBName      string              `yaml:"db_name"`
	SSLMode     string              `yaml:"ssl_mode"`
}

type PostgresCredentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
