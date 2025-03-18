package user

import (
	"awesomeProject2/handle"
	"github.com/gin-gonic/gin"
)

func Registry(r *gin.Engine) {
	userGroup := r.Group("/user").Use(handle.AuthCheck())
	{
		userGroup.POST("/query", handle.QuerySystemInfo)
		userGroup.POST("/insert", handle.InsertSystemUserInfo)
		userGroup.POST("/update", handle.UpdateSystemUserInfo)
		userGroup.POST("/delete", handle.DeleteSystemUserById)
	}

	userGroup = r.Group("/grpc").Use(handle.AuthCheck())
	{
		userGroup.POST("/test", handle.TestGrpcCall)
	}
}
