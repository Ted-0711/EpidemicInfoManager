package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MeasureRouter struct {
}

// InitMeasureRouter 初始化 Measure 路由信息
func (s *MeasureRouter) InitMeasureRouter(Router *gin.RouterGroup) {
	measureRouter := Router.Group("measure").Use(middleware.OperationRecord())
	measureRouterWithoutRecord := Router.Group("measure")
	var measureApi = v1.ApiGroupApp.EpidemicInfoApiGroup.MeasureApi
	{
		measureRouter.POST("createMeasure", measureApi.CreateMeasure)   // 新建Measure
		measureRouter.DELETE("deleteMeasure", measureApi.DeleteMeasure) // 删除Measure
		measureRouter.DELETE("deleteMeasureByIds", measureApi.DeleteMeasureByIds) // 批量删除Measure
		measureRouter.PUT("updateMeasure", measureApi.UpdateMeasure)    // 更新Measure
	}
	{
		measureRouterWithoutRecord.GET("findMeasure", measureApi.FindMeasure)        // 根据ID获取Measure
		measureRouterWithoutRecord.GET("getMeasureList", measureApi.GetMeasureList)  // 获取Measure列表
	}
}
