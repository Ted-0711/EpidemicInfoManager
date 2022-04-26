package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AreaRouter struct {
}

// InitAreaRouter 初始化 Area 路由信息
func (s *AreaRouter) InitAreaRouter(Router *gin.RouterGroup) {
	areaRouter := Router.Group("area").Use(middleware.OperationRecord())
	areaRouterWithoutRecord := Router.Group("area")
	var areaApi = v1.ApiGroupApp.EpidemicInfoApiGroup.AreaApi
	{
		areaRouter.POST("createArea", areaApi.CreateArea)   // 新建Area
		areaRouter.DELETE("deleteArea", areaApi.DeleteArea) // 删除Area
		areaRouter.DELETE("deleteAreaByIds", areaApi.DeleteAreaByIds) // 批量删除Area
		areaRouter.PUT("updateArea", areaApi.UpdateArea)    // 更新Area
	}
	{
		areaRouterWithoutRecord.GET("findArea", areaApi.FindArea)        // 根据ID获取Area
		areaRouterWithoutRecord.GET("getAreaList", areaApi.GetAreaList)  // 获取Area列表
	}
}
