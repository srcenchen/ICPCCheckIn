package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"player/internal/utils"
	"player/internal/utils/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type Device struct {
	Id       int
	Status   uint8
	Ip       string
	Mac      string
	StuName  string
	StuNum   string
	CheckIn  time.Time
	CheckOut time.Time
}

func Status(c *gin.Context) {
	// 获取status
	resp, err := http.Get(utils.Device.RemoteURL + "/v1/device/device-by-mac?mac=" + utils.Device.Mac)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务端错误"})
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	device := &Device{}
	_ = json.Unmarshal(body, &device)
	logger.Sugar().Info(device)
	if device.Id == -1 {
		// 去注册
		logger.Sugar().Info("没有注册")
		register(c)
		return
	}
	c.JSON(200, gin.H{
		"code":         device.Status, // 0 代表未签到 1代表已经签到 2代表已经签退
		"stuName":      device.StuName,
		"stuNum":       device.StuNum,
		"ip":           device.Ip,
		"mac":          device.Mac,
		"checkInTime":  device.CheckIn,
		"checkOutTime": device.CheckOut,
	})
}

func register(c *gin.Context) {
	b := map[string]string{
		"mac": utils.Device.Mac,
		"ip":  utils.Device.Address,
	}
	bj, _ := json.Marshal(b)
	logger.Sugar().Info(string(bj))
	resp, _ := http.Post(utils.Device.RemoteURL+"/v1/device/register", "application/json", bytes.NewReader(bj))
	body, _ := io.ReadAll(resp.Body)
	logger.Sugar().Info(string(body))
	Status(c)
}
