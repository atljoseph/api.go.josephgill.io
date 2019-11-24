package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3PublicAssetURL returns the URL of an asset in the public s3 bucket
// To do so, requires a relative path to the desired asset
func S3PublicAssetURL(relativePath string) string {
	funcTag := "S3PublicAssetURL"
	logMessage(funcTag, relativePath)
	return fmt.Sprintf("%s/%s/%s", config.S3PublicURL, config.S3PublicName, relativePath)
}

// S3PublicAssetList gets a list of the public assets in an aws s3 bucket
// TODO: Do more with aws sdk. List Objects, Rename Objects, Delete Objects, Upload New Objects, Upload Existing Objects
func S3PublicAssetList() {
	funcTag := "S3PublicAssetList"

	logMessage(funcTag, "get config and building request")
	cfg := aws.NewConfig().
		WithRegion(config.S3PublicRegion).
		WithCredentials(credentials.NewStaticCredentials(config.S3UserID, config.S3UserSecret, ""))
		// WithCredentials(credentials.AnonymousCredentials)
	svc := s3.New(session.New(cfg))
	input := &s3.ListObjectsInput{
		Bucket: aws.String(config.S3PublicName),
		// MaxKeys: aws.Int64(100),
	}

	logMessage(funcTag, "getting object list")
	result, err := svc.ListObjects(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				logError(funcTag, fmt.Errorf("%s: %s", s3.ErrCodeNoSuchBucket, aerr.Error()))
			default:
				logError(funcTag, aerr)
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			logError(funcTag, err)
		}
		return
	}

	logMessage(funcTag, fmt.Sprintf("found %d results", len(result.Contents)))
}
