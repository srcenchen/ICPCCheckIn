package info

import (
	"fmt"
	"os"
	"server/internal/data"
	"server/internal/data/model"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func Export(c *gin.Context) {
	_ = os.MkdirAll("./resource/tmp/", 0775)
	// 创建新的 Excel 文件
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// 创建工作表
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println("Error creating sheet:", err)
		return
	} // 设置表头（第一行）
	headers := []string{"序号", "姓名", "学号", "签到时间", "签退时间", "设备ip", "设备mac", "状态"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c%d", 'A'+i, 1)
		_ = f.SetCellValue("Sheet1", cell, header)
	}

	// 添加样式
	// 设置全局字体大小为11，居中对齐，允许文本换行
	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size: 11,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		fmt.Println("Error creating style:", err)
		return
	}
	// 获取所有devices
	var devices []model.Device
	data.DB().Find(&devices)
	// 应用样式到所有单元格（包括表头和数据行）
	for i := 0; i < len(headers); i++ {
		for j := 0; j < len(devices)+1; j++ { // 修正：去掉多余的+1
			cell, _ := excelize.CoordinatesToCellName(i+1, j+1)
			_ = f.SetCellStyle("Sheet1", cell, cell, style)
		}
	}

	//// 设置行高为57（包括表头和数据行）
	//for i := 1; i < len(devices)+1; i++ { // 修正：调整循环范围
	//	_ = f.SetRowHeight("Sheet1", i+1, 57)
	//}

	// 设置列宽
	for i := 0; i < len(headers); i++ {
		colName, _ := excelize.ColumnNumberToName(i + 1)
		_ = f.SetColWidth("Sheet1", colName, colName, float64(len(headers[i])+5)) // 基础宽度+5
	}

	// 调整特定列的宽度以适应内容
	_ = f.SetColWidth("Sheet1", "D", "D", 20)
	_ = f.SetColWidth("Sheet1", "E", "E", 20)

	// 设置图片列的宽度
	for i := 0; i < 10; i++ { // 假设最多10张图片
		colName, _ := excelize.ColumnNumberToName(11 + i)
		_ = f.SetColWidth("Sheet1", colName, colName, 20)
	}

	// 写入每条日志数据
	for i, log := range devices {
		row := i + 2 // 第二行开始写数据
		status := ""
		switch log.Status {
		case 0:
			status = "已注册"
			break
		case 1:
			status = "未签退"
			break
		case 2:
			status = "已签退"
			break
		}
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), i+1)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), log.StuName)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), log.StuNum)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), log.CheckIn)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), log.CheckOut)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), log.Ip)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), log.Mac)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("H%d", row), status)
	}
	// 设置默认工作表
	f.SetActiveSheet(index)

	// 保存文件
	err = f.SaveAs("./resource/tmp/export.xlsx")
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
	c.JSON(200, gin.H{
		"path": "resource/tmp/export.xlsx",
	})
}
