package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"player/internal/utils"
	"player/internal/utils/logger"
	"runtime"

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
	_ = ShutdownAfterDelay(60)
	c.JSON(200, gin.H{
		"code":    0, // 0 代表成功 1代表失败
		"message": "签退成功",
	})
}

// ShutdownAfterDelay 实现指定秒数后关机（适配多系统）
// delaySeconds: 延迟关机的秒数（建议按整分钟传值，适配Linux/macOS）
func ShutdownAfterDelay(delaySeconds int) error {
	var cmd *exec.Cmd
	osType := runtime.GOOS

	// 根据操作系统拼接对应的关机命令
	switch osType {
	case "windows":
		// Windows命令：shutdown /s /t 延迟秒数
		cmd = exec.Command("shutdown", "/s", "/t", fmt.Sprintf("%d", delaySeconds))
	case "linux", "darwin": // darwin是macOS的系统标识
		// Linux/macOS命令：shutdown -h +分钟数（60秒=1分钟）
		minutes := delaySeconds / 60
		cmd = exec.Command("shutdown", "-h", fmt.Sprintf("+%d", minutes))
	default:
		return fmt.Errorf("不支持的操作系统：%s", osType)
	}

	// 执行关机命令并捕获错误
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("执行关机命令失败：%v，输出信息：%s", err, string(output))
	}

	// 给用户清晰的提示
	fmt.Printf("✅ 已成功触发关机指令！\n")
	fmt.Printf("📌 系统类型：%s\n", osType)
	fmt.Printf("⏳ 电脑将在 %d 秒后自动关机！\n", delaySeconds)
	return nil
}
