package handle

import (
	"awesomeProject2/app"
	"awesomeProject2/middleware/logger"
	"awesomeProject2/middleware/redis"
	"awesomeProject2/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/rpc"
	"strconv"
	"time"
)

type User struct {
	Id    string `json:"id" binding:"required"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func QuerySystemInfo(c *gin.Context) {
	var user model.SystemUserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "绑定参数错误"})
		return
	}
	userStr, _ := redis.GetKey(strconv.Itoa(user.Id))
	if userStr != "" {
		json.Unmarshal([]byte(userStr), &user)
		logger.Info("get user from redis:", user)
		app.OK(c, user, "")
		return
	}

	//user = model.FindById(user.Id)
	client, err := rpc.Dial("tcp", "localhost:8082")

	defer client.Close()
	if err != nil {
		panic("客户端连接错误")
	}

	client.Call("User.GetUserInfo", user, &user)

	if user.Id != 0 {
		userStr, _ := json.Marshal(user)
		logger.Info("set user to redis:", string(userStr))
		redis.SetKey(strconv.Itoa(user.Id), string(userStr), 50*time.Second)
	}
	app.OK(c, user, "")

}

func InsertSystemUserInfo(c *gin.Context) {
	var user model.SystemUserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		app.ERROR(c, nil, "参数绑定错误", 400)
		return
	}
	id, insertE := model.Insert(user)
	if insertE != nil {
		app.ERROR(c, nil, insertE.Error(), 400)
		return
	}
	app.OK(c, id, "")
}

func UpdateSystemUserInfo(c *gin.Context) {
	var user model.SystemUserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		app.ERROR(c, nil, "参数绑定错误", 400)
		return
	}
	dbUser := model.FindById(user.Id)
	dbUser.UserName = user.UserName
	model.Update(dbUser)
	app.OK(c, dbUser, "success")
}

func DeleteSystemUserById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		app.ERROR(c, nil, "参数绑定错误", 400)
		return
	}
	idInt, _ := strconv.Atoi(id)
	dbUser := model.FindById(idInt)
	if dbUser.UserName == "" {
		app.ERROR(c, nil, "db用户不存在", 400)
		return
	}
	model.Delete(dbUser)
	app.OK(c, idInt, "")
}
