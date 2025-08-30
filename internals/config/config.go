package config

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

type DatabaseConfi struct {
	Host   string `koanf:"host" validate:"required"`
	Port   int `koanf:"port" validate:"required"`
	User   string `koanf:"user" validate:"required"`
	Password   int `koanf:"password"`
	Name   string `koanf:"name" validate:"required"`
	SSLMode   string `koanf:"ssl_mode" validate:"required"`
	MaxOpenConns  int `koanf:"max_open_conns" validate:"required"`
	MaxIdleConns  int `koanf:"max_idle_conns" validate:"required"`
	ConnMaxLifetime  int `koanf:"conn_max_lifetime" validate:"required"`
	ConnMaxIdletime  int `koanf:"conn_max_idle_time" validate:"required"`

}

type AuthConfig 

func LoadConfig(*ServerConfig, error){}