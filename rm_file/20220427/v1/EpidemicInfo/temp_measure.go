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

type Temp_measureApi struct {
}

var temp_measureService = service.ServiceGroupApp.EpidemicInfoServiceGroup.Temp_measureService


// CreateTemp_measure 创建Temp_measure
// @Tags Temp_measure
// @Summary 创建Temp_measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Temp_measure true "创建Temp_measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /temp_measure/createTemp_measure [post]
func (temp_measureApi *Temp_measureApi) CreateTemp_measure(c *gin.Context) {
	var temp_measure EpidemicInfo.Temp_measure
	_ = c.ShouldBindJSON(&temp_measure)
	if err := temp_measureService.CreateTemp_measure(temp_measure); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTemp_measure 删除Temp_measure
// @Tags Temp_measure
// @Summary 删除Temp_measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Temp_measure true "删除Temp_measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /temp_measure/deleteTemp_measure [delete]
func (temp_measureApi *Temp_measureApi) DeleteTemp_measure(c *gin.Context) {
	var temp_measure EpidemicInfo.Temp_measure
	_ = c.ShouldBindJSON(&temp_measure)
	if err := temp_measureService.DeleteTemp_measure(temp_measure); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTemp_measureByIds 批量删除Temp_measure
// @Tags Temp_measure
// @Summary 批量删除Temp_measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Temp_measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /temp_measure/deleteTemp_measureByIds [delete]
func (temp_measureApi *Temp_measureApi) DeleteTemp_measureByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := temp_measureService.DeleteTemp_measureByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTemp_measure 更新Temp_measure
// @Tags Temp_measure
// @Summary 更新Temp_measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Temp_measure true "更新Temp_measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /temp_measure/updateTemp_measure [put]
func (temp_measureApi *Temp_measureApi) UpdateTemp_measure(c *gin.Context) {
	var temp_measure EpidemicInfo.Temp_measure
	_ = c.ShouldBindJSON(&temp_measure)
	if err := temp_measureService.UpdateTemp_measure(temp_measure); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTemp_measure 用id查询Temp_measure
// @Tags Temp_measure
// @Summary 用id查询Temp_measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Temp_measure true "用id查询Temp_measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /temp_measure/findTemp_measure [get]
func (temp_measureApi *Temp_measureApi) FindTemp_measure(c *gin.Context) {
	var temp_measure EpidemicInfo.Temp_measure
	_ = c.ShouldBindQuery(&temp_measure)
	if err, retemp_measure := temp_measureService.GetTemp_measure(temp_measure.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retemp_measure": retemp_measure}, c)
	}
}

// GetTemp_measureList 分页获取Temp_measure列表
// @Tags Temp_measure
// @Summary 分页获取Temp_measure列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.Temp_measureSearch true "分页获取Temp_measure列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /temp_measure/getTemp_measureList [get]
func (temp_measureApi *Temp_measureApi) GetTemp_measureList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.Temp_measureSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := temp_measureService.GetTemp_measureInfoList(pageInfo); err != nil {
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
