package middleware

import "mime/multipart"

type Controller interface {
	UploadToAWS(string, string, multipart.File) error
	DownloadS3File(string, string) error
}
