package uploader

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Uploader() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials("AKIATCKAR3BXGNLV3ZZI", "A+EFWgoqNwgCHRrbOUwgV3dstE6FLntsW3aXrIzJ", ""),
	})
	if err != nil {
		log.Fatalf("failed to create session, %v", err)
	}
	uploader := s3manager.NewUploader(sess)
	f, err := os.Open("example1.txt")
	if err != nil {
		log.Fatalf("failed to open file  %v", err)
	}
	defer f.Close()

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("bucket-pradeep-test"),
		ACL:    aws.String("public-read"),
		Key:    aws.String("test1.txt"),
		Body:   f,
	})
	if err != nil {
		log.Fatalf("failed to upload file, %v", err)
	}

	fmt.Printf("Successfully uploaded  %q\n", result.Location)

}
func Downloader() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials("AKIATCKAR3BXGNLV3ZZI", "A+EFWgoqNwgCHRrbOUwgV3dstE6FLntsW3aXrIzJ", ""),
	})
	if err != nil {
		log.Fatalf("failed to create session, %v", err)
	}
	file, err := os.Create("dell.txt")
	if err != nil {
		log.Fatal(err)
	}
	down := s3manager.NewDownloader(sess)
	_, err = down.Download(file, &s3.GetObjectInput{
		Bucket: aws.String("bucket-pradeep-test"),
		Key:    aws.String("test1.txt"),
	})
	if err != nil {
		log.Fatalf("failed to download file, %v", err)
	}
	fmt.Printf("Successfully downloaded  ")
}
