package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ThermographRouter struct {
}

// InitThermographRouter 初始化 Thermograph 路由信息
func (s *ThermographRouter) InitThermographRouter(Router *gin.RouterGroup) {
	thermographRouter := Router.Group("thermograph").Use(middleware.OperationRecord())
	thermographRouterWithoutRecord := Router.Group("thermograph")
	var thermographApi = v1.ApiGroupApp.EpidemicInfoApiGroup.ThermographApi
	{
		thermographRouter.POST("createThermograph", thermographApi.CreateThermograph)   // 新建Thermograph
		thermographRouter.DELETE("deleteThermograph", thermographApi.DeleteThermograph) // 删除Thermograph
		thermographRouter.DELETE("deleteThermographByIds", thermographApi.DeleteThermographByIds) // 批量删除Thermograph
		thermographRouter.PUT("updateThermograph", thermographApi.UpdateThermograph)    // 更新Thermograph
	}
	{
		thermographRouterWithoutRecord.GET("findThermograph", thermographApi.FindThermograph)        // 根据ID获取Thermograph
		thermographRouterWithoutRecord.GET("getThermographList", thermographApi.GetThermographList)  // 获取Thermograph列表
	}
}
