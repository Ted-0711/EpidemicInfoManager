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

type CounselorApi struct {
}

var counselorService = service.ServiceGroupApp.EpidemicInfoServiceGroup.CounselorService


// CreateCounselor 创建Counselor
// @Tags Counselor
// @Summary 创建Counselor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Counselor true "创建Counselor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /counselor/createCounselor [post]
func (counselorApi *CounselorApi) CreateCounselor(c *gin.Context) {
	var counselor EpidemicInfo.Counselor
	_ = c.ShouldBindJSON(&counselor)
	if err := counselorService.CreateCounselor(counselor); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCounselor 删除Counselor
// @Tags Counselor
// @Summary 删除Counselor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Counselor true "删除Counselor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /counselor/deleteCounselor [delete]
func (counselorApi *CounselorApi) DeleteCounselor(c *gin.Context) {
	var counselor EpidemicInfo.Counselor
	_ = c.ShouldBindJSON(&counselor)
	if err := counselorService.DeleteCounselor(counselor); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCounselorByIds 批量删除Counselor
// @Tags Counselor
// @Summary 批量删除Counselor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Counselor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /counselor/deleteCounselorByIds [delete]
func (counselorApi *CounselorApi) DeleteCounselorByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := counselorService.DeleteCounselorByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCounselor 更新Counselor
// @Tags Counselor
// @Summary 更新Counselor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Counselor true "更新Counselor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /counselor/updateCounselor [put]
func (counselorApi *CounselorApi) UpdateCounselor(c *gin.Context) {
	var counselor EpidemicInfo.Counselor
	_ = c.ShouldBindJSON(&counselor)
	if err := counselorService.UpdateCounselor(counselor); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCounselor 用id查询Counselor
// @Tags Counselor
// @Summary 用id查询Counselor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Counselor true "用id查询Counselor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /counselor/findCounselor [get]
func (counselorApi *CounselorApi) FindCounselor(c *gin.Context) {
	var counselor EpidemicInfo.Counselor
	_ = c.ShouldBindQuery(&counselor)
	if err, recounselor := counselorService.GetCounselor(counselor.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recounselor": recounselor}, c)
	}
}

// GetCounselorList 分页获取Counselor列表
// @Tags Counselor
// @Summary 分页获取Counselor列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.CounselorSearch true "分页获取Counselor列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /counselor/getCounselorList [get]
func (counselorApi *CounselorApi) GetCounselorList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.CounselorSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := counselorService.GetCounselorInfoList(pageInfo); err != nil {
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
