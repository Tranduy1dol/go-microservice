package config

type Config struct {
	Server  ServerConfig
	MongoDB MongoConfig
	Redis   RedisConfig
	OAuth   OAuthConfig
}

type ServerConfig struct {
	Port string
	Env  string
}

type MongoConfig struct {
	URI      string
	Database string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       string
}

type OAuthConfig struct {
	GoogleClientID     string
	GoogleClientSecret string
	RedirectURL        string
	JWTSecret          string
}
