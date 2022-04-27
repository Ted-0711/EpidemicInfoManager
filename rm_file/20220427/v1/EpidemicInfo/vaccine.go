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

type VaccineApi struct {
}

var vaccineService = service.ServiceGroupApp.EpidemicInfoServiceGroup.VaccineService


// CreateVaccine 创建Vaccine
// @Tags Vaccine
// @Summary 创建Vaccine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Vaccine true "创建Vaccine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vaccine/createVaccine [post]
func (vaccineApi *VaccineApi) CreateVaccine(c *gin.Context) {
	var vaccine EpidemicInfo.Vaccine
	_ = c.ShouldBindJSON(&vaccine)
	if err := vaccineService.CreateVaccine(vaccine); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVaccine 删除Vaccine
// @Tags Vaccine
// @Summary 删除Vaccine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Vaccine true "删除Vaccine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vaccine/deleteVaccine [delete]
func (vaccineApi *VaccineApi) DeleteVaccine(c *gin.Context) {
	var vaccine EpidemicInfo.Vaccine
	_ = c.ShouldBindJSON(&vaccine)
	if err := vaccineService.DeleteVaccine(vaccine); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVaccineByIds 批量删除Vaccine
// @Tags Vaccine
// @Summary 批量删除Vaccine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Vaccine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /vaccine/deleteVaccineByIds [delete]
func (vaccineApi *VaccineApi) DeleteVaccineByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := vaccineService.DeleteVaccineByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVaccine 更新Vaccine
// @Tags Vaccine
// @Summary 更新Vaccine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Vaccine true "更新Vaccine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /vaccine/updateVaccine [put]
func (vaccineApi *VaccineApi) UpdateVaccine(c *gin.Context) {
	var vaccine EpidemicInfo.Vaccine
	_ = c.ShouldBindJSON(&vaccine)
	if err := vaccineService.UpdateVaccine(vaccine); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVaccine 用id查询Vaccine
// @Tags Vaccine
// @Summary 用id查询Vaccine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Vaccine true "用id查询Vaccine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /vaccine/findVaccine [get]
func (vaccineApi *VaccineApi) FindVaccine(c *gin.Context) {
	var vaccine EpidemicInfo.Vaccine
	_ = c.ShouldBindQuery(&vaccine)
	if err, revaccine := vaccineService.GetVaccine(vaccine.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revaccine": revaccine}, c)
	}
}

// GetVaccineList 分页获取Vaccine列表
// @Tags Vaccine
// @Summary 分页获取Vaccine列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.VaccineSearch true "分页获取Vaccine列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vaccine/getVaccineList [get]
func (vaccineApi *VaccineApi) GetVaccineList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.VaccineSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := vaccineService.GetVaccineInfoList(pageInfo); err != nil {
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
