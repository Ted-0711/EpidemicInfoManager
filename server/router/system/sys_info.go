package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InfoRouter struct{}

func (s *InfoRouter) GetWeatherRouter(Router *gin.RouterGroup) {
	weatherRouter := Router.Group("info").Use(middleware.OperationRecord())
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		weatherRouter.POST("getWeather", baseApi.GetWeather) // 管理员注册账号
	}
}
