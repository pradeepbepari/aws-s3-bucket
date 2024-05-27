package middleware

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsService interface {
	UploadToAws(string, string, multipart.File) (*s3manager.UploadOutput, error)
	DownloadFromAws(string, string) error
}
