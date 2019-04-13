package tcp_server

type Config struct {
	Port int
}

func (cfg Config) IsZero() bool {
	return cfg.Port == 0
}
