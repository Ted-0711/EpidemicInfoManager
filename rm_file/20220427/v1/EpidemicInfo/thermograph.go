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

type ThermographApi struct {
}

var thermographService = service.ServiceGroupApp.EpidemicInfoServiceGroup.ThermographService


// CreateThermograph 创建Thermograph
// @Tags Thermograph
// @Summary 创建Thermograph
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Thermograph true "创建Thermograph"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /thermograph/createThermograph [post]
func (thermographApi *ThermographApi) CreateThermograph(c *gin.Context) {
	var thermograph EpidemicInfo.Thermograph
	_ = c.ShouldBindJSON(&thermograph)
	if err := thermographService.CreateThermograph(thermograph); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteThermograph 删除Thermograph
// @Tags Thermograph
// @Summary 删除Thermograph
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Thermograph true "删除Thermograph"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /thermograph/deleteThermograph [delete]
func (thermographApi *ThermographApi) DeleteThermograph(c *gin.Context) {
	var thermograph EpidemicInfo.Thermograph
	_ = c.ShouldBindJSON(&thermograph)
	if err := thermographService.DeleteThermograph(thermograph); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteThermographByIds 批量删除Thermograph
// @Tags Thermograph
// @Summary 批量删除Thermograph
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Thermograph"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /thermograph/deleteThermographByIds [delete]
func (thermographApi *ThermographApi) DeleteThermographByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := thermographService.DeleteThermographByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateThermograph 更新Thermograph
// @Tags Thermograph
// @Summary 更新Thermograph
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Thermograph true "更新Thermograph"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /thermograph/updateThermograph [put]
func (thermographApi *ThermographApi) UpdateThermograph(c *gin.Context) {
	var thermograph EpidemicInfo.Thermograph
	_ = c.ShouldBindJSON(&thermograph)
	if err := thermographService.UpdateThermograph(thermograph); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindThermograph 用id查询Thermograph
// @Tags Thermograph
// @Summary 用id查询Thermograph
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Thermograph true "用id查询Thermograph"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /thermograph/findThermograph [get]
func (thermographApi *ThermographApi) FindThermograph(c *gin.Context) {
	var thermograph EpidemicInfo.Thermograph
	_ = c.ShouldBindQuery(&thermograph)
	if err, rethermograph := thermographService.GetThermograph(thermograph.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rethermograph": rethermograph}, c)
	}
}

// GetThermographList 分页获取Thermograph列表
// @Tags Thermograph
// @Summary 分页获取Thermograph列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.ThermographSearch true "分页获取Thermograph列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /thermograph/getThermographList [get]
func (thermographApi *ThermographApi) GetThermographList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.ThermographSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := thermographService.GetThermographInfoList(pageInfo); err != nil {
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
