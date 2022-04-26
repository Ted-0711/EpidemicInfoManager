package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Fill_inRouter struct {
}

// InitFill_inRouter 初始化 Fill_in 路由信息
func (s *Fill_inRouter) InitFill_inRouter(Router *gin.RouterGroup) {
	fill_inRouter := Router.Group("fill_in").Use(middleware.OperationRecord())
	fill_inRouterWithoutRecord := Router.Group("fill_in")
	var fill_inApi = v1.ApiGroupApp.EpidemicInfoApiGroup.Fill_inApi
	{
		fill_inRouter.POST("createFill_in", fill_inApi.CreateFill_in)   // 新建Fill_in
		fill_inRouter.DELETE("deleteFill_in", fill_inApi.DeleteFill_in) // 删除Fill_in
		fill_inRouter.DELETE("deleteFill_inByIds", fill_inApi.DeleteFill_inByIds) // 批量删除Fill_in
		fill_inRouter.PUT("updateFill_in", fill_inApi.UpdateFill_in)    // 更新Fill_in
	}
	{
		fill_inRouterWithoutRecord.GET("findFill_in", fill_inApi.FindFill_in)        // 根据ID获取Fill_in
		fill_inRouterWithoutRecord.GET("getFill_inList", fill_inApi.GetFill_inList)  // 获取Fill_in列表
	}
}
