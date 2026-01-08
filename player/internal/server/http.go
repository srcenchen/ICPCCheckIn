package server

import (
	"embed"
	"player/internal/service"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//go:embed public/*
var distFS embed.FS

// NewHTTPServer 创建HTTP服务
//
//go:em
func NewHTTPServer() {
	srv := gin.Default()
	// 静态资源
	f, _ := static.EmbedFolder(distFS, "public")
	srv.Use(static.Serve("/", f))
	v1 := srv.Group("/v1")
	v1.POST("/check-in", service.CheckIn)
	v1.POST("/check-out", service.CheckOut)
	v1.GET("/status", service.Status)
	srv.Run()
}
