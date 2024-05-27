package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	awscontroller "github.com/pradeepbepari/aws-cloud/aws_controller"
)

func RegisteredRoutes(routes *mux.Router, controler *awscontroller.AwsController) {
	routes.HandleFunc("/upload", controler.HandleUploadToAws).Methods(http.MethodPost)
	routes.HandleFunc("/download", controler.HandleDownloadsToAws).Methods(http.MethodGet)

}
