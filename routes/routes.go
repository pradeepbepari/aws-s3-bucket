package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	controllers "github.com/pradeep/aws/controllers"
)

func RegisterRoutes(router *mux.Router, controller *controllers.ControllerHandular) {
	router.HandleFunc("/upload", controller.UploadFileToAws).Methods(http.MethodPost)
	router.HandleFunc("/download", controller.DownloadFromAws).Methods(http.MethodGet)
}
