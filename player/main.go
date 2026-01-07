package main

import (
	"player/internal/server"
	"player/internal/utils"
	"player/internal/utils/logger"
)

func init() {
	_ = logger.Init() // 初始化 zap
	utils.Device = &utils.DeviceInfo{}
	err := utils.Device.GetDeviceMac()
	if err != nil {
		logger.Sugar().Error("获取设备信息出错:", err.Error())
		return
	}
	logger.Sugar().Info("设备信息：", utils.Device)
}

func main() {
	server.NewHTTPServer()
}
