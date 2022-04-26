package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Clock_inRouter struct {
}

// InitClock_inRouter 初始化 Clock_in 路由信息
func (s *Clock_inRouter) InitClock_inRouter(Router *gin.RouterGroup) {
	clock_inRouter := Router.Group("clock_in").Use(middleware.OperationRecord())
	clock_inRouterWithoutRecord := Router.Group("clock_in")
	var clock_inApi = v1.ApiGroupApp.EpidemicInfoApiGroup.Clock_inApi
	{
		clock_inRouter.POST("createClock_in", clock_inApi.CreateClock_in)   // 新建Clock_in
		clock_inRouter.DELETE("deleteClock_in", clock_inApi.DeleteClock_in) // 删除Clock_in
		clock_inRouter.DELETE("deleteClock_inByIds", clock_inApi.DeleteClock_inByIds) // 批量删除Clock_in
		clock_inRouter.PUT("updateClock_in", clock_inApi.UpdateClock_in)    // 更新Clock_in
	}
	{
		clock_inRouterWithoutRecord.GET("findClock_in", clock_inApi.FindClock_in)        // 根据ID获取Clock_in
		clock_inRouterWithoutRecord.GET("getClock_inList", clock_inApi.GetClock_inList)  // 获取Clock_in列表
	}
}
