package service

import "github.com/gin-gonic/gin"

func Status(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":         1, // 0 代表未签到 1代表已经签到 2代表已经签退
		"stuName":      "伞恩晨",
		"stuNum":       "25346121",
		"ip":           "192.168.1.2",
		"mac":          "E0:ES:ED:SA:XZ",
		"checkInTime":  "2026-1-1 06:00:00",
		"checkOutTime": "2026-1-1 07:00:00",
	})
}
