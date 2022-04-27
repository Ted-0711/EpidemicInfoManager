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

type ManufacturerApi struct {
}

var manufacturerService = service.ServiceGroupApp.EpidemicInfoServiceGroup.ManufacturerService


// CreateManufacturer 创建Manufacturer
// @Tags Manufacturer
// @Summary 创建Manufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Manufacturer true "创建Manufacturer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /manufacturer/createManufacturer [post]
func (manufacturerApi *ManufacturerApi) CreateManufacturer(c *gin.Context) {
	var manufacturer EpidemicInfo.Manufacturer
	_ = c.ShouldBindJSON(&manufacturer)
	if err := manufacturerService.CreateManufacturer(manufacturer); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteManufacturer 删除Manufacturer
// @Tags Manufacturer
// @Summary 删除Manufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Manufacturer true "删除Manufacturer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /manufacturer/deleteManufacturer [delete]
func (manufacturerApi *ManufacturerApi) DeleteManufacturer(c *gin.Context) {
	var manufacturer EpidemicInfo.Manufacturer
	_ = c.ShouldBindJSON(&manufacturer)
	if err := manufacturerService.DeleteManufacturer(manufacturer); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteManufacturerByIds 批量删除Manufacturer
// @Tags Manufacturer
// @Summary 批量删除Manufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Manufacturer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /manufacturer/deleteManufacturerByIds [delete]
func (manufacturerApi *ManufacturerApi) DeleteManufacturerByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := manufacturerService.DeleteManufacturerByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateManufacturer 更新Manufacturer
// @Tags Manufacturer
// @Summary 更新Manufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Manufacturer true "更新Manufacturer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /manufacturer/updateManufacturer [put]
func (manufacturerApi *ManufacturerApi) UpdateManufacturer(c *gin.Context) {
	var manufacturer EpidemicInfo.Manufacturer
	_ = c.ShouldBindJSON(&manufacturer)
	if err := manufacturerService.UpdateManufacturer(manufacturer); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindManufacturer 用id查询Manufacturer
// @Tags Manufacturer
// @Summary 用id查询Manufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Manufacturer true "用id查询Manufacturer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /manufacturer/findManufacturer [get]
func (manufacturerApi *ManufacturerApi) FindManufacturer(c *gin.Context) {
	var manufacturer EpidemicInfo.Manufacturer
	_ = c.ShouldBindQuery(&manufacturer)
	if err, remanufacturer := manufacturerService.GetManufacturer(manufacturer.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remanufacturer": remanufacturer}, c)
	}
}

// GetManufacturerList 分页获取Manufacturer列表
// @Tags Manufacturer
// @Summary 分页获取Manufacturer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.ManufacturerSearch true "分页获取Manufacturer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /manufacturer/getManufacturerList [get]
func (manufacturerApi *ManufacturerApi) GetManufacturerList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.ManufacturerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := manufacturerService.GetManufacturerInfoList(pageInfo); err != nil {
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
