package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Test_swabRouter struct {
}

// InitTest_swabRouter 初始化 Test_swab 路由信息
func (s *Test_swabRouter) InitTest_swabRouter(Router *gin.RouterGroup) {
	test_swabRouter := Router.Group("test_swab").Use(middleware.OperationRecord())
	test_swabRouterWithoutRecord := Router.Group("test_swab")
	var test_swabApi = v1.ApiGroupApp.EpidemicInfoApiGroup.Test_swabApi
	{
		test_swabRouter.POST("createTest_swab", test_swabApi.CreateTest_swab)   // 新建Test_swab
		test_swabRouter.DELETE("deleteTest_swab", test_swabApi.DeleteTest_swab) // 删除Test_swab
		test_swabRouter.DELETE("deleteTest_swabByIds", test_swabApi.DeleteTest_swabByIds) // 批量删除Test_swab
		test_swabRouter.PUT("updateTest_swab", test_swabApi.UpdateTest_swab)    // 更新Test_swab
	}
	{
		test_swabRouterWithoutRecord.GET("findTest_swab", test_swabApi.FindTest_swab)        // 根据ID获取Test_swab
		test_swabRouterWithoutRecord.GET("getTest_swabList", test_swabApi.GetTest_swabList)  // 获取Test_swab列表
	}
}
