package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"player/internal/utils"
	"player/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

type CheckInRequest struct {
	StuName string `json:"stuName" binding:"required"` // 选手姓名
	StuNum  string `json:"stuNum" binding:"required"`  // 选手学号
}

// CheckIn 签到接口
func CheckIn(c *gin.Context) {
	req := &CheckInRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "缺少参数",
		})
		return
	}
	b := map[string]string{
		"mac":     utils.Device.Mac,
		"stuName": req.StuName,
		"stuNum":  req.StuNum,
	}
	bj, _ := json.Marshal(b)
	logger.Sugar().Info(string(bj))
	resp, _ := http.Post(utils.Device.RemoteURL+"/v1/device/check-in", "application/json", bytes.NewReader(bj))
	body, _ := io.ReadAll(resp.Body)
	logger.Sugar().Info(string(body))
	c.JSON(200, gin.H{
		"code":    0, // 0 代表成功 1代表失败
		"message": req.StuName + req.StuNum + "签到成功",
	})
}

// CheckOut 签退接口
func CheckOut(c *gin.Context) {
	b := map[string]string{
		"mac": utils.Device.Mac,
	}
	bj, _ := json.Marshal(b)
	logger.Sugar().Info(string(bj))
	resp, _ := http.Post(utils.Device.RemoteURL+"/v1/device/check-out?mac="+utils.Device.Mac, "", nil)
	body, _ := io.ReadAll(resp.Body)
	logger.Sugar().Info(string(body))
	c.JSON(200, gin.H{
		"code":    0, // 0 代表成功 1代表失败
		"message": "签退成功",
	})
}
