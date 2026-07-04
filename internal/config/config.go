package config

type Config struct {
	Listen string
	Path   string
	Target string
	Debug  bool
}

func Default() *Config {
	return &Config{
		Listen: ":700",
		Path:   "/ssh-ws",
		Target: "127.0.0.1:22",
		Debug:  false,
	}
}
