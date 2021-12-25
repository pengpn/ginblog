package routes

import (
	v1 "gin-use/api/v1"
	"gin-use/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	route := r.Group("api/v1")
	{
		// user模块路由接口
		route.POST("user/add", v1.AddUser)
		route.GET("users", v1.GetUsers)
		route.PUT("user/:id", v1.EditUser)
		route.DELETE("user/:id", v1.DeleteUser)
		// category模块路由接口
		route.POST("category/add", v1.AddCategory)
		route.GET("categories", v1.GetCategories)
		route.PUT("category/:id", v1.EditCategory)
		route.DELETE("category/:id", v1.DeleteCategory)
		// article模块路由接口
	}

	r.Run(utils.HttpPort)
}
