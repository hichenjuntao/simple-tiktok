package internal

import (
	"fmt"
	"simple_tiktok_single/internal/dao/mysql"
	_ "simple_tiktok_single/internal/logic"
	"simple_tiktok_single/internal/router"
	"simple_tiktok_single/logs"
	"simple_tiktok_single/manifest/config"
	"simple_tiktok_single/pkg/snowflake"

	"go.uber.org/zap"
)

const ConfigFilePath = "manifest/config/config.yaml"

func Launch() {
	// 记载配置文件
	if err := config.Init(ConfigFilePath); err != nil {
		fmt.Printf("init config failed, Error:%v\n", err)
		return
	}

	// 加载日志配置
	if err := logs.Init(config.Conf.LogConfig); err != nil {
		fmt.Printf("init log failed, Error:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("log init success.")

	// 加载 mysql 配置
	if err := mysql.Init(config.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed, Error:%v\n", err)
		return
	}
	defer mysql.Close()

	// 加载 redis 配置
	// if err := redis.Init(config.Conf.RedisConfig); err != nil {
	// 	fmt.Printf("init redis failed, Error:%v\n", err)
	// 	return
	// }
	// defer redis.Close()

	// 加载雪花算法配置
	if err := snowflake.Init(config.Conf.StartTime, config.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, Error:%v\n", err)
		return
	}

	router.Setup(router.Init())
}
