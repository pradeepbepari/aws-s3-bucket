package awsmiddleware

import (
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pradeepbepari/aws-cloud/utils"
)

type AwsOperations struct {
}

func NewAwsOperation() *AwsOperations {
	return &AwsOperations{}
}
func (s *AwsOperations) UploadToAws(bucket, filekey string, file multipart.File) (*s3manager.UploadOutput, error) {
	sess, err := utils.AwsSession()
	if err != nil {
		return nil, err
	}
	uploader := s3manager.NewUploader(sess)
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filekey),
		ACL:    aws.String("public-read"),
		Body:   file,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *AwsOperations) DownloadFromAws(bucket, filekey string) error {
	sess, err := utils.AwsSession()
	if err != nil {
		return err
	}
	file, err := os.Create("/tem/upload.txt")
	if err != nil {
		return err
	}
	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filekey),
	})
	if err != nil {
		return err
	}
	return nil
}
