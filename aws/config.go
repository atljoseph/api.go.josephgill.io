package aws

import (
	"strings"
)

// Config is the configuration for this package
type Config struct {
	S3PublicName   string
	S3PublicURL    string
	S3PublicSecret string
}

// MergeWithDefaults merges the passed in config with the default options
func (cfg *Config) MergeWithDefaults() *Config {
	if strings.EqualFold(cfg.S3PublicName, "") {
		cfg.S3PublicName = "public.josephgill.io"
	}
	if strings.EqualFold(cfg.S3PublicURL, "") {
		cfg.S3PublicURL = "https://s3.us-east-2.amazonaws.com"
	}
	if strings.EqualFold(cfg.S3PublicSecret, "") {
		cfg.S3PublicSecret = "N/A"
	}
	return cfg
}
