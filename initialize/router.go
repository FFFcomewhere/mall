package initialize

import (
	v1 "exam8/router/v1"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	//无需验证权限的路由
	apiNotAuthV1 := r.Group("/api/v1")

	{
		//主页
		apiNotAuthV1.GET("/", func(context *gin.Context) {
			fmt.Println("hello world")
		})

		//用户登录
		apiNotAuthV1.POST("/login", v1.Login)
		//添加用户信息
		apiNotAuthV1.POST("/users", v1.AddUser)
		//查看用户信息
		apiNotAuthV1.GET("/users", v1.GetUser)
		//修改用户信息
		apiNotAuthV1.PUT("/users", v1.UpdateUser)
		//删除用户信息
		apiNotAuthV1.DELETE("/users", v1.DeleteUser)

	}

	return r
}
