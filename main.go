package main

import (
	"awesomeProject2/global"
	"awesomeProject2/middleware/logger"
	"awesomeProject2/middleware/redis"
	"awesomeProject2/route/user"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	name := flag.String("name", "wang", "用户名称")
	fmt.Println(name)
	flag.Parse()
	global.InitViper("settings.dev.yml", "./config")
	global.InitMysql(global.GetMysqlConfig())
	redis.InitRedis(global.GetRedisConfig())
	logger.Init()
	// 初始化 Gin 路由引擎（默认包含 Logger 和 Recovery 中间件）
	r := gin.Default()
	user.Registry(r)
	fmt.Println("start success")
	// 启动服务器（默认监听 0.0.0.0:8080
	port := global.Viper.GetString("settings.application.port")
	r.Run(fmt.Sprintf(":%s", port))
}
