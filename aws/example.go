package aws

// import (
// 	"fmt"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/awserr"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3"
// )

// func example() {
// 	svc := s3.New(session.New())
// 	input := &s3.ListObjectsInput{
// 		Bucket:  aws.String("examplebucket"),
// 		MaxKeys: aws.Int64(2),
// 	}

// 	result, err := svc.ListObjects(input)
// 	if err != nil {
// 		if aerr, ok := err.(awserr.Error); ok {
// 			switch aerr.Code() {
// 			case s3.ErrCodeNoSuchBucket:
// 				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
// 			default:
// 				fmt.Println(aerr.Error())
// 			}
// 		} else {
// 			// Print the error, cast err to awserr.Error to get the Code and
// 			// Message from an error.
// 			fmt.Println(err.Error())
// 		}
// 		return
// 	}

// 	fmt.Println(result)
// }
