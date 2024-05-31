package server

import (
	"encoding/json"
	"errors"
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
		var a args
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := s.s.LightOn(r.Context(), a.Name)
		handleErrorHttp(w, err)
	}
}

func (s *HttpServer) lightOff() http.HandlerFunc {
	type args struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var a args
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := s.s.LightOff(r.Context(), a.Name)
		handleErrorHttp(w, err)
	}
}

func (s *HttpServer) lightChangeColor() http.HandlerFunc {
	type args struct {
		Name  string        `json:"name"`
		Color *domain.Color `json:"color"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var a args
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := s.s.LightChangeColor(r.Context(), a.Name, a.Color)
		handleErrorHttp(w, err)
	}
}

func (s *HttpServer) lightChangeWhite() http.HandlerFunc {
	type args struct {
		Name  string        `json:"name"`
		White *domain.White `json:"white"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var a args
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := s.s.LightChangeWhite(r.Context(), a.Name, a.White)
		handleErrorHttp(w, err)
	}
}

// helpers

type response struct {
	Message string `json:"message"`
}

func newResponse(message string) response { return response{Message: message} }

func handleErrorHttp(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case errors.Is(err, domain.ErrorEventNotValid):
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(newResponse(err.Error()))
	case errors.Is(err, domain.ErrorDeviceNotFound):
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(newResponse(err.Error()))
	case err != nil:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(newResponse(err.Error()))
	default:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newResponse("success"))
	}

}
