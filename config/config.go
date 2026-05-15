package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server  ServerConfig
	MongoDB MongoConfig
	Redis   RedisConfig
	OAuth   OAuthConfig
	GRPC    GRPCConfig
}

type ServerConfig struct {
	Port          string `mapstructure:"port"`
	Env           string `mapstructure:"env"`
	LogLevel      string `mapstructure:"log_level"`
	EnableSwagger bool   `mapstructure:"enable_swagger"`
	UIBaseURL     string `mapstructure:"ui_base_url"`
}

type MongoConfig struct {
	URI      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type OAuthConfig struct {
	GoogleClientID     string `mapstructure:"google_client_id"`
	GoogleClientSecret string `mapstructure:"google_client_secret"`
	RedirectURL        string `mapstructure:"redirect_url"`
	JWTSecret          string `mapstructure:"jwt_secret"`
}

type GRPCConfig struct {
	SearchEngineAddr string `mapstructure:"search_engine_addr"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, relying on system environment variables")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("$HOME/.config/.kotoba-press")

	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.env", "development")
	viper.SetDefault("server.log_level", "info")
	viper.SetDefault("server.enable_swagger", false)
	viper.SetDefault("server.ui_base_url", "http://localhost:3000")
	viper.SetDefault("mongodb.uri", "mongodb://admin:secret@localhost:27017")
	viper.SetDefault("mongodb.database", "learning-japanese")
	viper.SetDefault("redis.addr", "localhost:6379")
	viper.SetDefault("redis.db", "0")
	viper.SetDefault("grpc.search_engine_addr", "")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
