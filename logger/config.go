package logger

// Config is the configuration for this package
type Config struct {
	// Filename string
}

// MergeWithDefaults merges the passed in config with the default options
func (cfg *Config) MergeWithDefaults() *Config {
	// if strings.EqualFold(cfg.Filename, "") {
	// 	cfg.Filename = ""
	// 	// cfg.Filename = "dev.log"
	// }
	return cfg
}
