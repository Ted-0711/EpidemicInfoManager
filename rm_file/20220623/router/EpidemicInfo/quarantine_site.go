package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Quarantine_siteRouter struct {
}

// InitQuarantine_siteRouter 初始化 Quarantine_site 路由信息
func (s *Quarantine_siteRouter) InitQuarantine_siteRouter(Router *gin.RouterGroup) {
	quarantine_siteRouter := Router.Group("quarantine_site").Use(middleware.OperationRecord())
	quarantine_siteRouterWithoutRecord := Router.Group("quarantine_site")
	var quarantine_siteApi = v1.ApiGroupApp.EpidemicInfoApiGroup.Quarantine_siteApi
	{
		quarantine_siteRouter.POST("createQuarantine_site", quarantine_siteApi.CreateQuarantine_site)   // 新建Quarantine_site
		quarantine_siteRouter.DELETE("deleteQuarantine_site", quarantine_siteApi.DeleteQuarantine_site) // 删除Quarantine_site
		quarantine_siteRouter.DELETE("deleteQuarantine_siteByIds", quarantine_siteApi.DeleteQuarantine_siteByIds) // 批量删除Quarantine_site
		quarantine_siteRouter.PUT("updateQuarantine_site", quarantine_siteApi.UpdateQuarantine_site)    // 更新Quarantine_site
	}
	{
		quarantine_siteRouterWithoutRecord.GET("findQuarantine_site", quarantine_siteApi.FindQuarantine_site)        // 根据ID获取Quarantine_site
		quarantine_siteRouterWithoutRecord.GET("getQuarantine_siteList", quarantine_siteApi.GetQuarantine_siteList)  // 获取Quarantine_site列表
	}
}
