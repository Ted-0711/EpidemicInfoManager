package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StudentRouter struct {
}

// InitStudentRouter 初始化 Student 路由信息
func (s *StudentRouter) InitStudentRouter(Router *gin.RouterGroup) {
	studentRouter := Router.Group("student").Use(middleware.OperationRecord())
	studentRouterWithoutRecord := Router.Group("student")
	var studentApi = v1.ApiGroupApp.EpidemicInfoApiGroup.StudentApi
	{
		studentRouter.POST("createStudent", studentApi.CreateStudent)   // 新建Student
		studentRouter.DELETE("deleteStudent", studentApi.DeleteStudent) // 删除Student
		studentRouter.DELETE("deleteStudentByIds", studentApi.DeleteStudentByIds) // 批量删除Student
		studentRouter.PUT("updateStudent", studentApi.UpdateStudent)    // 更新Student
	}
	{
		studentRouterWithoutRecord.GET("findStudent", studentApi.FindStudent)        // 根据ID获取Student
		studentRouterWithoutRecord.GET("getStudentList", studentApi.GetStudentList)  // 获取Student列表
	}
}
