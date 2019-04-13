package postgres

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func (cfg Config) IsZero() bool {
	return cfg.Port == 0
}
