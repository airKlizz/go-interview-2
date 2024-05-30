package server

import (
	"encoding/json"
	"net/http"

	"mynewgoproject/internal/core/domain"
	"mynewgoproject/internal/core/port/driving"
)

type HttpServer struct {
	s   driving.Server
	mux *http.ServeMux
}

func NewHttpServer(s driving.Server) *HttpServer {
	mux := http.NewServeMux()
	httpServer := &HttpServer{s: s, mux: mux}
	mux.HandleFunc("/light/on", httpServer.lightOn())
	mux.HandleFunc("/light/off", httpServer.lightOff())
	mux.HandleFunc("/light/color", httpServer.lightChangeColor())
	mux.HandleFunc("/light/white", httpServer.lightChangeWhite())
	return httpServer
}

func (s *HttpServer) Run() error {
	return http.ListenAndServe(":8080", s.mux)
}

func (s *HttpServer) lightOn() http.HandlerFunc {
	type args struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		panic("not implemented")
	}
}

func (s *HttpServer) lightOff() http.HandlerFunc {
	type args struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		panic("not implemented")
	}
}

func (s *HttpServer) lightChangeColor() http.HandlerFunc {
	type args struct {
		Name  string        `json:"name"`
		Color *domain.Color `json:"color"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		panic("not implemented")
	}
}

func (s *HttpServer) lightChangeWhite() http.HandlerFunc {
	type args struct {
		Name  string        `json:"name"`
		White *domain.White `json:"white"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		panic("not implemented")
	}
}

// helpers

type response struct {
	Message string `json:"message"`
}

func newResponse(message string) response { return response{Message: message} }

func handleErrorHttp(w http.ResponseWriter, err error) {
	// we cannot correctly handle error for the moment
	// it will be improved later
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(newResponse(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newResponse("success"))
	}
}
