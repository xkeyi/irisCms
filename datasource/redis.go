package datasource

import (
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"code/irisCms/config"
	"github.com/kataras/iris"
)

/**
 * 返回Redis实例
 */
func NewRedis() *redis.Database {
	var database *redis.Database

	// 项目配置
	cmsConfig := config.InitConfig()

	if cmsConfig != nil {
		rd := cmsConfig.Redis
		database = redis.New(redis.Config{
			Network:     rd.NetWork,
			Addr:        rd.Addr + ":" + rd.Port,
			Password:    rd.Password,
			Database:    "",
			MaxActive:   10,
			Prefix:      rd.Prefix,
		})
	} else {
		iris.New().Logger().Info(" config init error ")
	}

	return database
}