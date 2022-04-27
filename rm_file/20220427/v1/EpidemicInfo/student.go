package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type StudentApi struct {
}

var studentService = service.ServiceGroupApp.EpidemicInfoServiceGroup.StudentService


// CreateStudent 创建Student
// @Tags Student
// @Summary 创建Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Student true "创建Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/createStudent [post]
func (studentApi *StudentApi) CreateStudent(c *gin.Context) {
	var student EpidemicInfo.Student
	_ = c.ShouldBindJSON(&student)
	if err := studentService.CreateStudent(student); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStudent 删除Student
// @Tags Student
// @Summary 删除Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Student true "删除Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /student/deleteStudent [delete]
func (studentApi *StudentApi) DeleteStudent(c *gin.Context) {
	var student EpidemicInfo.Student
	_ = c.ShouldBindJSON(&student)
	if err := studentService.DeleteStudent(student); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStudentByIds 批量删除Student
// @Tags Student
// @Summary 批量删除Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /student/deleteStudentByIds [delete]
func (studentApi *StudentApi) DeleteStudentByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := studentService.DeleteStudentByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStudent 更新Student
// @Tags Student
// @Summary 更新Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Student true "更新Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /student/updateStudent [put]
func (studentApi *StudentApi) UpdateStudent(c *gin.Context) {
	var student EpidemicInfo.Student
	_ = c.ShouldBindJSON(&student)
	if err := studentService.UpdateStudent(student); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStudent 用id查询Student
// @Tags Student
// @Summary 用id查询Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Student true "用id查询Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /student/findStudent [get]
func (studentApi *StudentApi) FindStudent(c *gin.Context) {
	var student EpidemicInfo.Student
	_ = c.ShouldBindQuery(&student)
	if err, restudent := studentService.GetStudent(student.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restudent": restudent}, c)
	}
}

// GetStudentList 分页获取Student列表
// @Tags Student
// @Summary 分页获取Student列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.StudentSearch true "分页获取Student列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/getStudentList [get]
func (studentApi *StudentApi) GetStudentList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.StudentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := studentService.GetStudentInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
