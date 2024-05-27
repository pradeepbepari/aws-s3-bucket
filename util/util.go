package util

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func AwsS3Session() (*session.Session, error) {
	endpoint := os.Getenv("S3_ENDPOINT")

	if endpoint == "" {
		endpoint = "http://localhost:4566"
	}
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(os.Getenv("AWS_DEFAULT_REGION")),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(endpoint),
		Credentials:      credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		log.Fatalf("failed to create session, %v", err)
	}
	return sess, nil
}
