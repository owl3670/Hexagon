package config

type Config struct {
	JWTSecret string
	Port      int
	Postgres  PostgresConfig
	Redis     RedisConfig
}

type PostgresConfig struct {
	DSN         string
	MaxOpenConn int
	MaxIdleConn int
	MaxIdleTime string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}
