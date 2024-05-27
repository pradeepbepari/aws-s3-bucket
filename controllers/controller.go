package testaws

import (
	"log"
	"net/http"
	"os"

	middleware "github.com/pradeep/aws/middleware"
)

type ControllerHandular struct {
	controllers middleware.Controller
}

func NewHandular(control middleware.Controller) *ControllerHandular {
	return &ControllerHandular{controllers: control}

}
func (s *ControllerHandular) UploadFileToAws(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(50 << 50)
	if err != nil {
		log.Printf("failed to parse form %v", err)
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	bucket := os.Getenv("AWS_BUCKET_NAME")
	filename := header.Filename
	err = s.controllers.UploadToAWS(bucket, filename, file)
	if err != nil {
		http.Error(w, "Unable to upload file", http.StatusInternalServerError)
		return
	}
}
func (s *ControllerHandular) DownloadFromAws(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50 << 50)
	if err != nil {
		log.Printf("failed to parse form %v", err)
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	bucket := os.Getenv("AWS_BUCKET_NAME")
	filename := header.Filename

	err = s.controllers.DownloadS3File(bucket, filename)
	if err != nil {
		log.Fatal(err)
	}
}
