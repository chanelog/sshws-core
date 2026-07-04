package config

type Config struct {
	Listen  string
	Path    string
	Backend string
	Debug   bool
}

func Default() *Config {
	return &Config{
		Listen:  ":700",
		Path:    "/ssh-ws",
		Backend: "127.0.0.1:22",
		Debug:   false,
	}
}
