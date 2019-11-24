package aws

import (
	"strings"
)

// Config is the configuration for this package
type Config struct {
	S3PublicRegion string
	S3PublicName   string
	S3PublicURL    string
	S3UserID       string
	S3UserSecret   string
}

// MergeWithDefaults merges the passed in config with the default options
func (cfg *Config) MergeWithDefaults() *Config {
	if strings.EqualFold(cfg.S3PublicRegion, "") {
		cfg.S3PublicRegion = "us-east-2"
	}
	if strings.EqualFold(cfg.S3PublicName, "") {
		cfg.S3PublicName = "public.josephgill.io"
	}
	if strings.EqualFold(cfg.S3PublicURL, "") {
		cfg.S3PublicURL = "https://s3.us-east-2.amazonaws.com"
	}
	// READ ONLY USER:
	if strings.EqualFold(cfg.S3UserID, "") {
		cfg.S3UserID = "AKIA5A6MNRYQMQHB7THS"
	}
	if strings.EqualFold(cfg.S3UserSecret, "") {
		cfg.S3UserSecret = "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"
	}
	return cfg
}
