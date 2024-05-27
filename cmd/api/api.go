package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	controllers "github.com/pradeep/aws/controllers"
	upload "github.com/pradeep/aws/middleware/upload"
	"github.com/pradeep/aws/routes"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{addr: addr}
}
func (s *ApiServer) Run() error {
	backend := upload.NewBackend()
	controller := controllers.NewHandular(backend)
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	routes.RegisterRoutes(subrouter, controller)
	log.Printf("connection successful on port :%s", os.Getenv("PORT"))
	return http.ListenAndServe(s.addr, router)
}
