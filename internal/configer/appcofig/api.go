package appcofig

type API struct {
	ListenAddr string                `yaml:"listen_addr"`
	Log        Logger                `yaml:"log"`
	BasePath   string                `yaml:"base_path"`
	DocsPath   string                `yaml:"docs_path"`
	Postgres   PostgresConfiguration `yaml:"postgres"`
	Redis      RedisConfiguration    `yaml:"redis"`
	JWT        JWTConfiguration      `yaml:"jwt"`
}
