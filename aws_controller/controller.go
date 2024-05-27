package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pradeepbepari/aws-cloud/middleware"
	"github.com/pradeepbepari/aws-cloud/utils"
)

type AwsController struct {
	service middleware.AwsService
}

func NewAwsService(service middleware.AwsService) *AwsController {
	return &AwsController{service: service}
}
func (s *AwsController) HandleUploadToAws(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method invalid header request"))
		return
	}
	err := r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("max file size should be 10mb"))
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to get file from header"))
		return
	}
	defer file.Close()
	filekey := header.Filename
	bucket := os.Getenv("AWS_BUCKET")
	uploadurl, err := s.service.UploadToAws(bucket, filekey, file)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to upload file in s3 "))
		return
	}
	utils.WriteJson(w, http.StatusOK, uploadurl.Location)
}
func (s *AwsController) HandleDownloadsToAws(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf(" method invalid header request"))
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to get file from header"))
		return
	}
	defer file.Close()
	filekey := header.Filename
	bucket := os.Getenv("AWS_BUCKET")
	err = s.service.DownloadFromAws(bucket, filekey)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to download file from s3"))
		return
	}
	utils.WriteJson(w, http.StatusOK, "file downloaded")
}
