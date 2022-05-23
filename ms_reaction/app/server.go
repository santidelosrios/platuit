package app

import (
	"net/http"

	"github.com/santidelosrios/platuit/ms_reaction/app/handler"
	"github.com/santidelosrios/platuit/ms_reaction/cmd"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

//Server - struct for Project service server
type Server struct {
	configuration *cmd.Config
	router        *chi.Mux
	handler       *handler.Handler
}

//Close - closes the connection to server
func (s *Server) Close() {

}

//Start - starts a new server
func (s *Server) Start() {
	logrus.Infof("Starting Project service HTTP server on %v", ":"+s.configuration.Port)

	err := http.ListenAndServe(":"+s.configuration.Port, s.router)

	if err != nil {
		logrus.WithError(err).Fatal("error starting project service HTTP server")
	}
}

//NewServer - function to create a new server
func NewServer(config *cmd.Config, handler *handler.Handler) *Server {
	return &Server{configuration: config, handler: handler}
}

//SetupRoutes - function that setups all the routes of the service
func (s *Server) SetupRoutes() {
	s.router = chi.NewRouter()

	s.router.Route("/reaction", func(r chi.Router) {
		r.Post("/visit", s.handler.CreateVisitToTuit)
		//r.Post("/like")
	})
}
