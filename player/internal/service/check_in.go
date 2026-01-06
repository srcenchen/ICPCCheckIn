package service

import "github.com/gin-gonic/gin"

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
	c.JSON(200, gin.H{
		"code":    0, // 0 代表成功 1代表失败
		"message": req.StuName + req.StuNum + "签到成功",
	})
}

// CheckOut 签退接口
func CheckOut(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    0, // 0 代表成功 1代表失败
		"message": "签退成功",
	})
}
