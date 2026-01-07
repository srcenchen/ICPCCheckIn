package device

import (
	"errors"
	"net/http"
	"server/internal/data"
	"server/internal/data/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterReq struct {
	Mac string `json:"mac"`
	Ip  string `json:"ip"`
}

// Register 设备启动时候注册设备
func Register(c *gin.Context) {
	req := &RegisterReq{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数传输错误"})
	}
	device := model.Device{
		Status: 0,
		Ip:     req.Ip,
		Mac:    req.Mac,
	}
	// 删掉对应的 mac 地址
	data.DB().Where("mac = ?", req.Mac).Delete(&model.Device{})
	data.DB().Create(&device)
	c.JSON(http.StatusOK, gin.H{"message": "设备注册成功"})
}

// GetAllDevices 获取全部的设备
func GetAllDevices(c *gin.Context) {
	var devices []model.Device
	data.DB().Find(&devices)
	c.JSON(http.StatusOK, gin.H{"devices": devices})
}

// GetDeviceByMac 根据 mac 获取设备
func GetDeviceByMac(c *gin.Context) {
	device := &model.Device{}
	mac := c.Query("mac")
	err := data.DB().Where("mac = ?", mac).First(&device).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, model.Device{Id: -1})
		return
	}
	c.JSON(http.StatusOK, device)
}

type CheckInRequest struct {
	Mac     string `json:"mac"`                        // mac 地址
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
	device := model.Device{}
	data.DB().Where("mac = ?", req.Mac).First(&device)
	device.Status = 1
	device.StuNum = req.StuNum
	device.StuName = req.StuName
	device.CheckIn = time.Now()
	data.DB().Save(&device)
	c.JSON(200, gin.H{
		"code":    0, // 0 代表成功 1代表失败
		"message": req.StuName + req.StuNum + "签到成功",
	})
}

// CheckOut 签退接口
func CheckOut(c *gin.Context) {
	var device model.Device
	data.DB().Where("mac = ?", c.Query("mac")).First(&device)
	device.Status = 2
	device.CheckOut = time.Now()
	data.DB().Save(&device)
	c.JSON(200, gin.H{
		"code":    0, // 0 代表成功 1代表失败
		"message": "签退成功",
	})
}

// Delete 删除接口
func Delete(c *gin.Context) {
	if c.Query("type") != "" {
		// 全部删除
		data.DB().Delete(&model.Device{})
		c.JSON(200, gin.H{
			"code":    0, // 0 代表成功 1代表失败
			"message": "删除成功",
		})
		return
	}
	data.DB().Where("mac = ?", c.Query("mac")).Delete(&model.Device{})
	c.JSON(200, gin.H{
		"code":    0, // 0 代表成功 1代表失败
		"message": "删除成功",
	})
}
