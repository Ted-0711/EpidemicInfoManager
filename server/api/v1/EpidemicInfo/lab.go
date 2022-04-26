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

type LabApi struct {
}

var labService = service.ServiceGroupApp.EpidemicInfoServiceGroup.LabService


// CreateLab 创建Lab
// @Tags Lab
// @Summary 创建Lab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Lab true "创建Lab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lab/createLab [post]
func (labApi *LabApi) CreateLab(c *gin.Context) {
	var lab EpidemicInfo.Lab
	_ = c.ShouldBindJSON(&lab)
	if err := labService.CreateLab(lab); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLab 删除Lab
// @Tags Lab
// @Summary 删除Lab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Lab true "删除Lab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lab/deleteLab [delete]
func (labApi *LabApi) DeleteLab(c *gin.Context) {
	var lab EpidemicInfo.Lab
	_ = c.ShouldBindJSON(&lab)
	if err := labService.DeleteLab(lab); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLabByIds 批量删除Lab
// @Tags Lab
// @Summary 批量删除Lab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Lab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lab/deleteLabByIds [delete]
func (labApi *LabApi) DeleteLabByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := labService.DeleteLabByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLab 更新Lab
// @Tags Lab
// @Summary 更新Lab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Lab true "更新Lab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lab/updateLab [put]
func (labApi *LabApi) UpdateLab(c *gin.Context) {
	var lab EpidemicInfo.Lab
	_ = c.ShouldBindJSON(&lab)
	if err := labService.UpdateLab(lab); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLab 用id查询Lab
// @Tags Lab
// @Summary 用id查询Lab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Lab true "用id查询Lab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lab/findLab [get]
func (labApi *LabApi) FindLab(c *gin.Context) {
	var lab EpidemicInfo.Lab
	_ = c.ShouldBindQuery(&lab)
	if err, relab := labService.GetLab(lab.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relab": relab}, c)
	}
}

// GetLabList 分页获取Lab列表
// @Tags Lab
// @Summary 分页获取Lab列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.LabSearch true "分页获取Lab列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lab/getLabList [get]
func (labApi *LabApi) GetLabList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.LabSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := labService.GetLabInfoList(pageInfo); err != nil {
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
