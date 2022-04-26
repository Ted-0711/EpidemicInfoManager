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

type Test_swabApi struct {
}

var test_swabService = service.ServiceGroupApp.EpidemicInfoServiceGroup.Test_swabService


// CreateTest_swab 创建Test_swab
// @Tags Test_swab
// @Summary 创建Test_swab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Test_swab true "创建Test_swab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test_swab/createTest_swab [post]
func (test_swabApi *Test_swabApi) CreateTest_swab(c *gin.Context) {
	var test_swab EpidemicInfo.Test_swab
	_ = c.ShouldBindJSON(&test_swab)
	if err := test_swabService.CreateTest_swab(test_swab); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTest_swab 删除Test_swab
// @Tags Test_swab
// @Summary 删除Test_swab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Test_swab true "删除Test_swab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test_swab/deleteTest_swab [delete]
func (test_swabApi *Test_swabApi) DeleteTest_swab(c *gin.Context) {
	var test_swab EpidemicInfo.Test_swab
	_ = c.ShouldBindJSON(&test_swab)
	if err := test_swabService.DeleteTest_swab(test_swab); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTest_swabByIds 批量删除Test_swab
// @Tags Test_swab
// @Summary 批量删除Test_swab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Test_swab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /test_swab/deleteTest_swabByIds [delete]
func (test_swabApi *Test_swabApi) DeleteTest_swabByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := test_swabService.DeleteTest_swabByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTest_swab 更新Test_swab
// @Tags Test_swab
// @Summary 更新Test_swab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Test_swab true "更新Test_swab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /test_swab/updateTest_swab [put]
func (test_swabApi *Test_swabApi) UpdateTest_swab(c *gin.Context) {
	var test_swab EpidemicInfo.Test_swab
	_ = c.ShouldBindJSON(&test_swab)
	if err := test_swabService.UpdateTest_swab(test_swab); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTest_swab 用id查询Test_swab
// @Tags Test_swab
// @Summary 用id查询Test_swab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Test_swab true "用id查询Test_swab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /test_swab/findTest_swab [get]
func (test_swabApi *Test_swabApi) FindTest_swab(c *gin.Context) {
	var test_swab EpidemicInfo.Test_swab
	_ = c.ShouldBindQuery(&test_swab)
	if err, retest_swab := test_swabService.GetTest_swab(test_swab.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retest_swab": retest_swab}, c)
	}
}

// GetTest_swabList 分页获取Test_swab列表
// @Tags Test_swab
// @Summary 分页获取Test_swab列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.Test_swabSearch true "分页获取Test_swab列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test_swab/getTest_swabList [get]
func (test_swabApi *Test_swabApi) GetTest_swabList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.Test_swabSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := test_swabService.GetTest_swabInfoList(pageInfo); err != nil {
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
