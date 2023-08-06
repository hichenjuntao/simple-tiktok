package mysql

import (
	"fmt"
	"simple_tiktok_single/manifest/config"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	engine *gorm.DB
	once   = sync.Once{}
)

func Init(cfg *config.MysqlConfig) (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName)
	engine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // 开启预编译功能
	})
	if err != nil {
		zap.L().Error("gorm new engine failed, Error:", zap.Error(err))
		return
	}

	sqlDB, err := engine.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	return
}

func Close() (err error) {
	return
}
