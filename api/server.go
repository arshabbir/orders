package api

import (
	"fmt"
	"net/http"

	"github.com/arshabbir/utils/config"
	"github.com/arshabbir/utils/httpclient"
	"github.com/arshabbir/utils/logger"
	"github.com/gorilla/mux"
)

type server struct {
	m          *mux.Router
	l          logger.Logger
	conf       *config.Config
	httpClient httpclient.HttpClient
}

type Server interface {
	Start() error
}

func NewOrderServer(conf *config.Config, l logger.Logger) Server {
	return &server{m: mux.NewRouter(), l: l, conf: conf, httpClient: httpclient.NewHttpClient(conf, l)}
}

func (s *server) Start() error {

	// Register the order service end points
	s.m.HandleFunc("/ping", s.handlePing).Methods("GET")
	s.m.HandleFunc("/order", s.handleOrder).Methods("POST")

	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.conf.AppPort), s.m); err != nil {
		return err
	}

	return nil
}
