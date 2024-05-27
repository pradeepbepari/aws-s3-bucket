package uploader

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pradeep/aws/util"
)

type Backend struct {
}

func NewBackend() *Backend {
	return &Backend{}
}

func (s *Backend) UploadToAWS(bucket, filename string, file multipart.File) error {
	sess, err := util.AwsS3Session()
	if err != nil {
		log.Printf("failed to create aws session %+v", err)
		return nil
	}
	uploader := s3manager.NewUploader(sess)
	// svc := s3.New(sess)
	// path, err := svc.PutObject(&s3.PutObjectInput{
	// 	ACL:    aws.String("public-read"),
	// 	Bucket: aws.String(bucket),
	// 	Key:    aws.String(filename),
	// 	Body:   file,
	// })
	res, err := uploader.Upload(&s3manager.UploadInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		log.Printf("failed to upload in aws  %+v", err)
	}
	//log.Printf("successfully uploaded to s3  %s", path)
	log.Printf("successfully uploaded to s3  %q", res.Location)
	return nil
}
func (s *Backend) DownloadS3File(bucket, fileKey string) error {
	sess, err := util.AwsS3Session()
	if err != nil {
		log.Printf("failed to create aws session %+v", err)

	}
	file, err := os.Create("/tmp/upload.go")
	if err != nil {
		log.Println("Failed to create file", err)
	}
	defer file.Close()
	downloader := s3manager.NewDownloader(sess)
	path, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileKey),
	})
	if err != nil {
		log.Printf("failed to download file  %+v", err)

	}
	log.Printf("successfully downloaded to s3  %v", path)
	return nil
}
