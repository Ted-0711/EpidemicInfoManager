package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CounselorRouter struct {
}

// InitCounselorRouter 初始化 Counselor 路由信息
func (s *CounselorRouter) InitCounselorRouter(Router *gin.RouterGroup) {
	counselorRouter := Router.Group("counselor").Use(middleware.OperationRecord())
	counselorRouterWithoutRecord := Router.Group("counselor")
	var counselorApi = v1.ApiGroupApp.EpidemicInfoApiGroup.CounselorApi
	{
		counselorRouter.POST("createCounselor", counselorApi.CreateCounselor)   // 新建Counselor
		counselorRouter.DELETE("deleteCounselor", counselorApi.DeleteCounselor) // 删除Counselor
		counselorRouter.DELETE("deleteCounselorByIds", counselorApi.DeleteCounselorByIds) // 批量删除Counselor
		counselorRouter.PUT("updateCounselor", counselorApi.UpdateCounselor)    // 更新Counselor
	}
	{
		counselorRouterWithoutRecord.GET("findCounselor", counselorApi.FindCounselor)        // 根据ID获取Counselor
		counselorRouterWithoutRecord.GET("getCounselorList", counselorApi.GetCounselorList)  // 获取Counselor列表
	}
}
