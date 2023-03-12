package router

import (
	"github/IM_gin+vue/docs"
	"github/IM_gin+vue/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//首页
	r.GET("/index", service.GetIndex)

	//用户模块
	r.GET("/user/getUserList", service.GetUserList)
	return r
}
