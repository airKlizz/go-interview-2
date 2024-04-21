package server

import (
	"net/http"

	"mynewgoproject/internal/core/domain"
	"mynewgoproject/internal/core/port/driving"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	s driving.Server
	r *gin.Engine
}

func NewHttpServer(s driving.Server) *HttpServer {
	r := gin.Default()
	httpServer := &HttpServer{r: r, s: s}
	lightR := r.Group("/light")
	lightR.POST("/on", httpServer.lightOn())
	lightR.POST("/off", httpServer.lightOff())
	lightR.POST("/color", httpServer.lightChangeColor())
	lightR.POST("/white", httpServer.lightChangeWhite())
	return httpServer
}

func (s *HttpServer) Run() error {
	return s.r.Run()
}

func (s *HttpServer) lightOn() func(ctx *gin.Context) {
	type args struct {
		Name string
	}
	return func(ctx *gin.Context) {
		panic("not implemented")
	}
}

func (s *HttpServer) lightOff() func(ctx *gin.Context) {
	type args struct {
		Name string
	}
	return func(ctx *gin.Context) {
		panic("not implemented")
	}
}

func (s *HttpServer) lightChangeColor() func(ctx *gin.Context) {
	type args struct {
		Name  string
		Color *domain.Color
	}
	return func(ctx *gin.Context) {
		panic("not implemented")
	}
}

func (s *HttpServer) lightChangeWhite() func(ctx *gin.Context) {
	type args struct {
		Name  string
		White *domain.White
	}
	return func(ctx *gin.Context) {
		panic("not implemented")
	}
}

// helpers

type response struct {
	Message string `json:"message"`
}

func newResponse(message string) response { return response{Message: message} }

func handleErrorHttp(ctx *gin.Context, err error) {
	// we cannot correctly handle error for the moment
	// it will be improved later
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newResponse(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, newResponse("success"))
	}
}
