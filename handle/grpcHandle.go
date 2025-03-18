package handle

import (
	"awesomeProject2/app"
	"awesomeProject2/grpc/client"
	"github.com/gin-gonic/gin"
	"log"
)

func TestGrpcCall(c *gin.Context) {
	res := client.TestCallRemoteGrpc()
	log.Println(res)
	app.OK(c, res, "")
}
