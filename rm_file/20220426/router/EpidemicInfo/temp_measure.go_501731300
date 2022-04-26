package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Temp_measureRouter struct {
}

// InitTemp_measureRouter 初始化 Temp_measure 路由信息
func (s *Temp_measureRouter) InitTemp_measureRouter(Router *gin.RouterGroup) {
	temp_measureRouter := Router.Group("temp_measure").Use(middleware.OperationRecord())
	temp_measureRouterWithoutRecord := Router.Group("temp_measure")
	var temp_measureApi = v1.ApiGroupApp.EpidemicInfoApiGroup.Temp_measureApi
	{
		temp_measureRouter.POST("createTemp_measure", temp_measureApi.CreateTemp_measure)   // 新建Temp_measure
		temp_measureRouter.DELETE("deleteTemp_measure", temp_measureApi.DeleteTemp_measure) // 删除Temp_measure
		temp_measureRouter.DELETE("deleteTemp_measureByIds", temp_measureApi.DeleteTemp_measureByIds) // 批量删除Temp_measure
		temp_measureRouter.PUT("updateTemp_measure", temp_measureApi.UpdateTemp_measure)    // 更新Temp_measure
	}
	{
		temp_measureRouterWithoutRecord.GET("findTemp_measure", temp_measureApi.FindTemp_measure)        // 根据ID获取Temp_measure
		temp_measureRouterWithoutRecord.GET("getTemp_measureList", temp_measureApi.GetTemp_measureList)  // 获取Temp_measure列表
	}
}
