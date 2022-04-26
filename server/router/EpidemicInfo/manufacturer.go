package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ManufacturerRouter struct {
}

// InitManufacturerRouter 初始化 Manufacturer 路由信息
func (s *ManufacturerRouter) InitManufacturerRouter(Router *gin.RouterGroup) {
	manufacturerRouter := Router.Group("manufacturer").Use(middleware.OperationRecord())
	manufacturerRouterWithoutRecord := Router.Group("manufacturer")
	var manufacturerApi = v1.ApiGroupApp.EpidemicInfoApiGroup.ManufacturerApi
	{
		manufacturerRouter.POST("createManufacturer", manufacturerApi.CreateManufacturer)   // 新建Manufacturer
		manufacturerRouter.DELETE("deleteManufacturer", manufacturerApi.DeleteManufacturer) // 删除Manufacturer
		manufacturerRouter.DELETE("deleteManufacturerByIds", manufacturerApi.DeleteManufacturerByIds) // 批量删除Manufacturer
		manufacturerRouter.PUT("updateManufacturer", manufacturerApi.UpdateManufacturer)    // 更新Manufacturer
	}
	{
		manufacturerRouterWithoutRecord.GET("findManufacturer", manufacturerApi.FindManufacturer)        // 根据ID获取Manufacturer
		manufacturerRouterWithoutRecord.GET("getManufacturerList", manufacturerApi.GetManufacturerList)  // 获取Manufacturer列表
	}
}
