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

type MeasureApi struct {
}

var measureService = service.ServiceGroupApp.EpidemicInfoServiceGroup.MeasureService


// CreateMeasure 创建Measure
// @Tags Measure
// @Summary 创建Measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Measure true "创建Measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /measure/createMeasure [post]
func (measureApi *MeasureApi) CreateMeasure(c *gin.Context) {
	var measure EpidemicInfo.Measure
	_ = c.ShouldBindJSON(&measure)
	if err := measureService.CreateMeasure(measure); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMeasure 删除Measure
// @Tags Measure
// @Summary 删除Measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Measure true "删除Measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /measure/deleteMeasure [delete]
func (measureApi *MeasureApi) DeleteMeasure(c *gin.Context) {
	var measure EpidemicInfo.Measure
	_ = c.ShouldBindJSON(&measure)
	if err := measureService.DeleteMeasure(measure); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMeasureByIds 批量删除Measure
// @Tags Measure
// @Summary 批量删除Measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /measure/deleteMeasureByIds [delete]
func (measureApi *MeasureApi) DeleteMeasureByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := measureService.DeleteMeasureByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMeasure 更新Measure
// @Tags Measure
// @Summary 更新Measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Measure true "更新Measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /measure/updateMeasure [put]
func (measureApi *MeasureApi) UpdateMeasure(c *gin.Context) {
	var measure EpidemicInfo.Measure
	_ = c.ShouldBindJSON(&measure)
	if err := measureService.UpdateMeasure(measure); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMeasure 用id查询Measure
// @Tags Measure
// @Summary 用id查询Measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Measure true "用id查询Measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /measure/findMeasure [get]
func (measureApi *MeasureApi) FindMeasure(c *gin.Context) {
	var measure EpidemicInfo.Measure
	_ = c.ShouldBindQuery(&measure)
	if err, remeasure := measureService.GetMeasure(measure.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remeasure": remeasure}, c)
	}
}

// GetMeasureList 分页获取Measure列表
// @Tags Measure
// @Summary 分页获取Measure列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.MeasureSearch true "分页获取Measure列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /measure/getMeasureList [get]
func (measureApi *MeasureApi) GetMeasureList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.MeasureSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := measureService.GetMeasureInfoList(pageInfo); err != nil {
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
