package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	aws_controler "github.com/pradeepbepari/aws-cloud/aws_controller"
	awsmiddleware "github.com/pradeepbepari/aws-cloud/middleware/aws_middleware"
	"github.com/pradeepbepari/aws-cloud/routes"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{addr: addr}
}
func (s *ApiServer) Run() error {

	aws_service := awsmiddleware.NewAwsOperation()
	aws_controler := aws_controler.NewAwsService(aws_service)
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v2").Subrouter()
	routes.RegisteredRoutes(subrouter, aws_controler)
	log.Printf("Http sever listening on : %s", os.Getenv("PORT"))
	return http.ListenAndServe(s.addr, router)
}
