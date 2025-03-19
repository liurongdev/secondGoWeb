package main

import (
	"awesomeProject2/global"
	"awesomeProject2/middleware/logger"
	"awesomeProject2/middleware/redis"
	"awesomeProject2/route/user"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	start()
	//testChan()
}

func start() {
	// 初始化 Gin 路由引擎（默认包含 Logger 和 Recovery 中间件）
	r := gin.Default()
	user.Registry(r)
	fmt.Println("start success")
	// 启动服务器（默认监听 0.0.0.0:8080
	port := global.Viper.GetString("settings.application.port")
	r.Run(fmt.Sprintf(":%s", port))
}

func init() {
	name := flag.String("name", "wang", "用户名称")
	fmt.Println(name)
	flag.Parse()
	global.InitViper("settings.dev.yml", "./config")
	global.InitMysql(global.GetMysqlConfig())
	redis.InitRedis(global.GetRedisConfig())
	logger.Init()
}

func testChan() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received2:", msg2)
	default:
		fmt.Println("No reveived:")
		time.Sleep(500 * time.Microsecond)
	}

}
