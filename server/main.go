package main

import (
	"os"
	"server/internal/data"
	"server/internal/server"
	"server/internal/utils/logger"
)

func IsDirExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
func main() {
	if !IsDirExists("./resource") {
		_ = os.Mkdir("./resource", os.ModePerm)
	}
	_ = logger.Init() // 初始化 zap
	data.Init("./resource/data.db")
	server.NewHTTPServer()
}
