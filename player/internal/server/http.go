package server

import (
	"player/internal/service"

	"github.com/gin-gonic/gin"
)

// NewHTTPServer 创建HTTP服务
func NewHTTPServer() {
	srv := gin.Default()
	v1 := srv.Group("/v1")
	v1.POST("/check-in", service.CheckIn)
	v1.POST("/check-out", service.CheckOut)
	v1.GET("status", service.Status)
	srv.Run()
}
