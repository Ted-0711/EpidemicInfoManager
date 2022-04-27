package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LabRouter struct {
}

// InitLabRouter 初始化 Lab 路由信息
func (s *LabRouter) InitLabRouter(Router *gin.RouterGroup) {
	labRouter := Router.Group("lab").Use(middleware.OperationRecord())
	labRouterWithoutRecord := Router.Group("lab")
	var labApi = v1.ApiGroupApp.EpidemicInfoApiGroup.LabApi
	{
		labRouter.POST("createLab", labApi.CreateLab)   // 新建Lab
		labRouter.DELETE("deleteLab", labApi.DeleteLab) // 删除Lab
		labRouter.DELETE("deleteLabByIds", labApi.DeleteLabByIds) // 批量删除Lab
		labRouter.PUT("updateLab", labApi.UpdateLab)    // 更新Lab
	}
	{
		labRouterWithoutRecord.GET("findLab", labApi.FindLab)        // 根据ID获取Lab
		labRouterWithoutRecord.GET("getLabList", labApi.GetLabList)  // 获取Lab列表
	}
}
