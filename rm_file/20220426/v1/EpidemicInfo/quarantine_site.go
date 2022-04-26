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

type Quarantine_siteApi struct {
}

var quarantine_siteService = service.ServiceGroupApp.EpidemicInfoServiceGroup.Quarantine_siteService


// CreateQuarantine_site 创建Quarantine_site
// @Tags Quarantine_site
// @Summary 创建Quarantine_site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Quarantine_site true "创建Quarantine_site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /quarantine_site/createQuarantine_site [post]
func (quarantine_siteApi *Quarantine_siteApi) CreateQuarantine_site(c *gin.Context) {
	var quarantine_site EpidemicInfo.Quarantine_site
	_ = c.ShouldBindJSON(&quarantine_site)
	if err := quarantine_siteService.CreateQuarantine_site(quarantine_site); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuarantine_site 删除Quarantine_site
// @Tags Quarantine_site
// @Summary 删除Quarantine_site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Quarantine_site true "删除Quarantine_site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /quarantine_site/deleteQuarantine_site [delete]
func (quarantine_siteApi *Quarantine_siteApi) DeleteQuarantine_site(c *gin.Context) {
	var quarantine_site EpidemicInfo.Quarantine_site
	_ = c.ShouldBindJSON(&quarantine_site)
	if err := quarantine_siteService.DeleteQuarantine_site(quarantine_site); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuarantine_siteByIds 批量删除Quarantine_site
// @Tags Quarantine_site
// @Summary 批量删除Quarantine_site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Quarantine_site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /quarantine_site/deleteQuarantine_siteByIds [delete]
func (quarantine_siteApi *Quarantine_siteApi) DeleteQuarantine_siteByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := quarantine_siteService.DeleteQuarantine_siteByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuarantine_site 更新Quarantine_site
// @Tags Quarantine_site
// @Summary 更新Quarantine_site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Quarantine_site true "更新Quarantine_site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /quarantine_site/updateQuarantine_site [put]
func (quarantine_siteApi *Quarantine_siteApi) UpdateQuarantine_site(c *gin.Context) {
	var quarantine_site EpidemicInfo.Quarantine_site
	_ = c.ShouldBindJSON(&quarantine_site)
	if err := quarantine_siteService.UpdateQuarantine_site(quarantine_site); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuarantine_site 用id查询Quarantine_site
// @Tags Quarantine_site
// @Summary 用id查询Quarantine_site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Quarantine_site true "用id查询Quarantine_site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /quarantine_site/findQuarantine_site [get]
func (quarantine_siteApi *Quarantine_siteApi) FindQuarantine_site(c *gin.Context) {
	var quarantine_site EpidemicInfo.Quarantine_site
	_ = c.ShouldBindQuery(&quarantine_site)
	if err, requarantine_site := quarantine_siteService.GetQuarantine_site(quarantine_site.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requarantine_site": requarantine_site}, c)
	}
}

// GetQuarantine_siteList 分页获取Quarantine_site列表
// @Tags Quarantine_site
// @Summary 分页获取Quarantine_site列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.Quarantine_siteSearch true "分页获取Quarantine_site列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /quarantine_site/getQuarantine_siteList [get]
func (quarantine_siteApi *Quarantine_siteApi) GetQuarantine_siteList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.Quarantine_siteSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := quarantine_siteService.GetQuarantine_siteInfoList(pageInfo); err != nil {
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
