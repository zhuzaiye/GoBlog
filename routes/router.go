// File:    router
// Version: 1.0.0
// Creator: JoeLang
// Date:    2020/8/30 21:45
// DESC:

package routes

import (
	v1 "GoBlog/api/v1"
	"GoBlog/utils/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(setting.AppMode)
	engine := gin.Default()

	routerV1 := engine.Group("api/v1")
	{
		//Login Router

		//User Router
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("user/rows", v1.SearchUserList)
		routerV1.PUT("user/edit", v1.EditUserInfo)
		routerV1.DELETE("user/delete", v1.DeleteUser)

		//Article Router

		//Category Router
	}

	_ = engine.Run(setting.HttpPort)
}
