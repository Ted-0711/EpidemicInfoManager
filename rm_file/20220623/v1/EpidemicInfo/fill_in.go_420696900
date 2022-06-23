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

type Fill_inApi struct {
}

var fill_inService = service.ServiceGroupApp.EpidemicInfoServiceGroup.Fill_inService


// CreateFill_in 创建Fill_in
// @Tags Fill_in
// @Summary 创建Fill_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Fill_in true "创建Fill_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fill_in/createFill_in [post]
func (fill_inApi *Fill_inApi) CreateFill_in(c *gin.Context) {
	var fill_in EpidemicInfo.Fill_in
	_ = c.ShouldBindJSON(&fill_in)
	if err := fill_inService.CreateFill_in(fill_in); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFill_in 删除Fill_in
// @Tags Fill_in
// @Summary 删除Fill_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Fill_in true "删除Fill_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fill_in/deleteFill_in [delete]
func (fill_inApi *Fill_inApi) DeleteFill_in(c *gin.Context) {
	var fill_in EpidemicInfo.Fill_in
	_ = c.ShouldBindJSON(&fill_in)
	if err := fill_inService.DeleteFill_in(fill_in); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFill_inByIds 批量删除Fill_in
// @Tags Fill_in
// @Summary 批量删除Fill_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Fill_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fill_in/deleteFill_inByIds [delete]
func (fill_inApi *Fill_inApi) DeleteFill_inByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := fill_inService.DeleteFill_inByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFill_in 更新Fill_in
// @Tags Fill_in
// @Summary 更新Fill_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Fill_in true "更新Fill_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fill_in/updateFill_in [put]
func (fill_inApi *Fill_inApi) UpdateFill_in(c *gin.Context) {
	var fill_in EpidemicInfo.Fill_in
	_ = c.ShouldBindJSON(&fill_in)
	if err := fill_inService.UpdateFill_in(fill_in); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFill_in 用id查询Fill_in
// @Tags Fill_in
// @Summary 用id查询Fill_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Fill_in true "用id查询Fill_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fill_in/findFill_in [get]
func (fill_inApi *Fill_inApi) FindFill_in(c *gin.Context) {
	var fill_in EpidemicInfo.Fill_in
	_ = c.ShouldBindQuery(&fill_in)
	if err, refill_in := fill_inService.GetFill_in(fill_in.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refill_in": refill_in}, c)
	}
}

// GetFill_inList 分页获取Fill_in列表
// @Tags Fill_in
// @Summary 分页获取Fill_in列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.Fill_inSearch true "分页获取Fill_in列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fill_in/getFill_inList [get]
func (fill_inApi *Fill_inApi) GetFill_inList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.Fill_inSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := fill_inService.GetFill_inInfoList(pageInfo); err != nil {
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
