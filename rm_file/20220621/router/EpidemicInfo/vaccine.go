package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VaccineRouter struct {
}

// InitVaccineRouter 初始化 Vaccine 路由信息
func (s *VaccineRouter) InitVaccineRouter(Router *gin.RouterGroup) {
	vaccineRouter := Router.Group("vaccine").Use(middleware.OperationRecord())
	vaccineRouterWithoutRecord := Router.Group("vaccine")
	var vaccineApi = v1.ApiGroupApp.EpidemicInfoApiGroup.VaccineApi
	{
		vaccineRouter.POST("createVaccine", vaccineApi.CreateVaccine)   // 新建Vaccine
		vaccineRouter.DELETE("deleteVaccine", vaccineApi.DeleteVaccine) // 删除Vaccine
		vaccineRouter.DELETE("deleteVaccineByIds", vaccineApi.DeleteVaccineByIds) // 批量删除Vaccine
		vaccineRouter.PUT("updateVaccine", vaccineApi.UpdateVaccine)    // 更新Vaccine
	}
	{
		vaccineRouterWithoutRecord.GET("findVaccine", vaccineApi.FindVaccine)        // 根据ID获取Vaccine
		vaccineRouterWithoutRecord.GET("getVaccineList", vaccineApi.GetVaccineList)  // 获取Vaccine列表
	}
}
