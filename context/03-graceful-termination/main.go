package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type HttpServer struct {
	server http.Server
	mux    *http.ServeMux
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		server: http.Server{
			Addr: fmt.Sprintf("%s:%d", "localhost", 8080),
		},
		mux: http.NewServeMux(),
	}
}

func (s *HttpServer) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(pattern, handler)
}

func (s *HttpServer) Start(ctx context.Context) error {
	s.server.Handler = s.mux
	go func() { s.server.ListenAndServe() }()

	<-ctx.Done()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		if err := s.server.Close(); err != nil {
			fmt.Println("Could not close the server")
			return err
		}
	}

	return nil
}
