package server

import (
	"errors"
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
		var a args
		if err := ctx.ShouldBindJSON(&a); err != nil {
			ctx.JSON(http.StatusBadRequest, newResponse(err.Error()))
			return
		}
		err := s.s.LightOn(ctx, a.Name)
		handleErrorHttp(ctx, err)
	}
}

func (s *HttpServer) lightOff() func(ctx *gin.Context) {
	type args struct {
		Name string
	}
	return func(ctx *gin.Context) {
		var a args
		if err := ctx.ShouldBindJSON(&a); err != nil {
			ctx.JSON(http.StatusBadRequest, newResponse(err.Error()))
			return
		}
		err := s.s.LightOff(ctx, a.Name)
		handleErrorHttp(ctx, err)
	}
}

func (s *HttpServer) lightChangeColor() func(ctx *gin.Context) {
	type args struct {
		Name  string
		Color *domain.Color
	}
	return func(ctx *gin.Context) {
		var a args
		if err := ctx.ShouldBindJSON(&a); err != nil {
			ctx.JSON(http.StatusBadRequest, newResponse(err.Error()))
			return
		}
		err := s.s.LightChangeColor(ctx, a.Name, a.Color)
		handleErrorHttp(ctx, err)
	}
}

func (s *HttpServer) lightChangeWhite() func(ctx *gin.Context) {
	type args struct {
		Name  string
		White *domain.White
	}
	return func(ctx *gin.Context) {
		var a args
		if err := ctx.ShouldBindJSON(&a); err != nil {
			ctx.JSON(http.StatusBadRequest, newResponse(err.Error()))
			return
		}
		err := s.s.LightChangeWhite(ctx, a.Name, a.White)
		handleErrorHttp(ctx, err)
	}
}

// helpers

type response struct {
	Message string `json:"message"`
}

func newResponse(message string) response { return response{Message: message} }

func handleErrorHttp(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrorEventNotValid):
		ctx.JSON(http.StatusBadRequest, newResponse(err.Error()))
	case errors.Is(err, domain.ErrorDeviceNotFound):
		ctx.JSON(http.StatusNotFound, newResponse(err.Error()))
	case err != nil:
		ctx.JSON(http.StatusInternalServerError, newResponse(err.Error()))
	default:
		ctx.JSON(http.StatusOK, newResponse("success"))
	}
}
