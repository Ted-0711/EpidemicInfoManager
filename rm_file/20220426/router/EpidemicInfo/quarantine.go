package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QuarantineRouter struct {
}

// InitQuarantineRouter 初始化 Quarantine 路由信息
func (s *QuarantineRouter) InitQuarantineRouter(Router *gin.RouterGroup) {
	quarantineRouter := Router.Group("quarantine").Use(middleware.OperationRecord())
	quarantineRouterWithoutRecord := Router.Group("quarantine")
	var quarantineApi = v1.ApiGroupApp.EpidemicInfoApiGroup.QuarantineApi
	{
		quarantineRouter.POST("createQuarantine", quarantineApi.CreateQuarantine)   // 新建Quarantine
		quarantineRouter.DELETE("deleteQuarantine", quarantineApi.DeleteQuarantine) // 删除Quarantine
		quarantineRouter.DELETE("deleteQuarantineByIds", quarantineApi.DeleteQuarantineByIds) // 批量删除Quarantine
		quarantineRouter.PUT("updateQuarantine", quarantineApi.UpdateQuarantine)    // 更新Quarantine
	}
	{
		quarantineRouterWithoutRecord.GET("findQuarantine", quarantineApi.FindQuarantine)        // 根据ID获取Quarantine
		quarantineRouterWithoutRecord.GET("getQuarantineList", quarantineApi.GetQuarantineList)  // 获取Quarantine列表
	}
}
