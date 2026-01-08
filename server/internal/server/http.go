package server

import (
	"embed"
	d1 "server/internal/service/device"
	"server/internal/service/info"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//go:embed public/*
var distFS embed.FS

// NewHTTPServer 创建HTTP服务
func NewHTTPServer() {
	srv := gin.Default()
	v1 := srv.Group("/v1")
	// 静态资源
	f, _ := static.EmbedFolder(distFS, "public")
	srv.Use(static.Serve("/", f))
	// 挂载"./resource"到"/resource"，并且开启浏览器访问
	srv.Use(static.Serve("/resource", static.LocalFile("./resource", true)))
	// 设备组
	device := v1.Group("/device")
	device.POST("/register", d1.Register)           // 设备上线注册
	device.GET("/all-devices", d1.GetAllDevices)    // 获取全部设备
	device.GET("/device-by-mac", d1.GetDeviceByMac) // 根据 mac 获取设备信息
	device.POST("/check-in", d1.CheckIn)            // 签到
	device.POST("/check-out", d1.CheckOut)          // 签退
	device.POST("/delete", d1.Delete)               // 删除
	v1.GET("/export", info.Export)
	srv.Run(":8081")
}
