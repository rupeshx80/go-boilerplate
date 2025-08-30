package config

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)

// parent struct
type Config struct {
	Primary  Primary        `konf:"primary" validate:"required"`
	Server   ServerConfig   `konf:"server" validate:"required"`
	Database DatabaseConfig `konf:"database" validate:"required"`
	Auth     AuthConfig     `konf:"auth" validate:"required"`
	Redis    RedisConfig    `konf:"redis" validate:"required"`
}

type Primary struct {
	Env string `koanf: "env" validate:"required"`
}

type ServerConfig struct {
	Port               string   `konf:"port" validate:"required"`
	ReadTimeout        int      `konf:"read_timeout" validate:"required"`
	WriteTimeout       int      `konf:"write_timeout" validate:"required"`
	IdleTimeout        int      `konf:"idle_timeout" validate:"required"`
	CORSAllowedOrigins []string `konf:"cors_allowed_origins" validate:"required"`
}

type RedisConfig struct {
	Address string `koanf:"address" validate:"required"`
}

type DatabaseConfig struct {
	Host            string `koanf:"host" validate:"required"`
	Port            int    `koanf:"port" validate:"required"`
	User            string `koanf:"user" validate:"required"`
	Password        int    `koanf:"password"`
	Name            string `koanf:"name" validate:"required"`
	SSLMode         string `koanf:"ssl_mode" validate:"required"`
	MaxOpenConns    int    `koanf:"max_open_conns" validate:"required"`
	MaxIdleConns    int    `koanf:"max_idle_conns" validate:"required"`
	ConnMaxLifetime int    `koanf:"conn_max_lifetime" validate:"required"`
	ConnMaxIdletime int    `koanf:"conn_max_idle_time" validate:"required"`
}

type AuthConfig struct {
	SecretKey string `koanf:"secret_key" validate:"required"`
}

type Config struct {
	Env string
}

func LoadConfig() (*Config, error) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	k := koanf.New(".")

	err := k.Load(env.Provider("BOILERPLATE_", ".", func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s,"BOILERPLATE_"))
	}), nil)
	if(err !=nil){
		logger.Fatal().Err(err).Msg("could not load initial env variables")
	}
	
	mainConfig := &Config{}
	// fmt.Println("checking whats in &config", &Config)
	
	err = k.Unmarshal("", mainConfig)
	
	if err != nil {
		logger.Fatal().Err(err).Msg("could not unmarshal main config")
		
	}
}
