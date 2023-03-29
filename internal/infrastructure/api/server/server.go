package server

import (
	"OwnersAnimals/internal/config"
	"OwnersAnimals/internal/entities/animal"
	"OwnersAnimals/internal/entities/owner"
	"context"
	"fmt"
	"net/http"
	"time"
)

type Service interface {
	ReadOwnerByID(id int) (owner.Owner, error)
	ReadAnimals(ids []int) (animals []animal.Animal, err error)
}

type Server struct {
	srv  http.Server
	serv Service
}

func NewServer(conf config.API, h http.Handler) *Server {
	addr := fmt.Sprintf("%s:%d", conf.APIHost, conf.APIPort)

	s := &Server{}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       time.Duration(conf.APIReadTimeOut) * time.Second,
		WriteTimeout:      time.Duration(conf.APIWriteTimeOut) * time.Second,
		ReadHeaderTimeout: time.Duration(conf.APIReadHeadTimeOut) * time.Second,
	}
	return s
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	s.srv.Shutdown(ctx)
	cancel()
}
func (s *Server) Start(eap Service) {
	s.serv = eap
	go s.srv.ListenAndServe()
}
