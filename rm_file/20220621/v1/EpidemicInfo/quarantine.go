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

type QuarantineApi struct {
}

var quarantineService = service.ServiceGroupApp.EpidemicInfoServiceGroup.QuarantineService


// CreateQuarantine 创建Quarantine
// @Tags Quarantine
// @Summary 创建Quarantine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Quarantine true "创建Quarantine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /quarantine/createQuarantine [post]
func (quarantineApi *QuarantineApi) CreateQuarantine(c *gin.Context) {
	var quarantine EpidemicInfo.Quarantine
	_ = c.ShouldBindJSON(&quarantine)
	if err := quarantineService.CreateQuarantine(quarantine); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuarantine 删除Quarantine
// @Tags Quarantine
// @Summary 删除Quarantine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Quarantine true "删除Quarantine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /quarantine/deleteQuarantine [delete]
func (quarantineApi *QuarantineApi) DeleteQuarantine(c *gin.Context) {
	var quarantine EpidemicInfo.Quarantine
	_ = c.ShouldBindJSON(&quarantine)
	if err := quarantineService.DeleteQuarantine(quarantine); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuarantineByIds 批量删除Quarantine
// @Tags Quarantine
// @Summary 批量删除Quarantine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Quarantine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /quarantine/deleteQuarantineByIds [delete]
func (quarantineApi *QuarantineApi) DeleteQuarantineByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := quarantineService.DeleteQuarantineByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuarantine 更新Quarantine
// @Tags Quarantine
// @Summary 更新Quarantine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Quarantine true "更新Quarantine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /quarantine/updateQuarantine [put]
func (quarantineApi *QuarantineApi) UpdateQuarantine(c *gin.Context) {
	var quarantine EpidemicInfo.Quarantine
	_ = c.ShouldBindJSON(&quarantine)
	if err := quarantineService.UpdateQuarantine(quarantine); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuarantine 用id查询Quarantine
// @Tags Quarantine
// @Summary 用id查询Quarantine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Quarantine true "用id查询Quarantine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /quarantine/findQuarantine [get]
func (quarantineApi *QuarantineApi) FindQuarantine(c *gin.Context) {
	var quarantine EpidemicInfo.Quarantine
	_ = c.ShouldBindQuery(&quarantine)
	if err, requarantine := quarantineService.GetQuarantine(quarantine.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requarantine": requarantine}, c)
	}
}

// GetQuarantineList 分页获取Quarantine列表
// @Tags Quarantine
// @Summary 分页获取Quarantine列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.QuarantineSearch true "分页获取Quarantine列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /quarantine/getQuarantineList [get]
func (quarantineApi *QuarantineApi) GetQuarantineList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.QuarantineSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := quarantineService.GetQuarantineInfoList(pageInfo); err != nil {
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
