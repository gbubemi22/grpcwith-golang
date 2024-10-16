package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vendorService/internal/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(gin.Logger())

	r.GET("/", s.HelloWorldHandler)

	r.Use(middleware.ErrorHandlerMiddleware)
	r.NoRoute(middleware.HandleNotFound)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}
