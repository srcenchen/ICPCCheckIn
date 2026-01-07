package data

import (
	"server/internal/data/model"
	"server/internal/utils/logger"
	"sync"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Init(dbPath string) {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err != nil {
			logger.Sugar().Fatal("数据库初始化失败", err)
		}
		// 合并表
		err = db.AutoMigrate(&model.Device{})
		if err != nil {
			logger.Sugar().Fatal("数据库合并失败", err)
		}
	})
}

func DB() *gorm.DB {
	return db
}
