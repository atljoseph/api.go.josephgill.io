package aws

import "fmt"

// S3PublicAssetURL returns the URL of an asset in the public s3 bucket
// To do so, requires a relative path to the desired asset
func S3PublicAssetURL(relativePath string) string {
	return fmt.Sprintf("%s/%s/%s", config.S3PublicURL, config.S3PublicName, relativePath)
}
