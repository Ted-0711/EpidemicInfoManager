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

type Clock_inApi struct {
}

var clock_inService = service.ServiceGroupApp.EpidemicInfoServiceGroup.Clock_inService


// CreateClock_in 创建Clock_in
// @Tags Clock_in
// @Summary 创建Clock_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Clock_in true "创建Clock_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clock_in/createClock_in [post]
func (clock_inApi *Clock_inApi) CreateClock_in(c *gin.Context) {
	var clock_in EpidemicInfo.Clock_in
	_ = c.ShouldBindJSON(&clock_in)
	if err := clock_inService.CreateClock_in(clock_in); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteClock_in 删除Clock_in
// @Tags Clock_in
// @Summary 删除Clock_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Clock_in true "删除Clock_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clock_in/deleteClock_in [delete]
func (clock_inApi *Clock_inApi) DeleteClock_in(c *gin.Context) {
	var clock_in EpidemicInfo.Clock_in
	_ = c.ShouldBindJSON(&clock_in)
	if err := clock_inService.DeleteClock_in(clock_in); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteClock_inByIds 批量删除Clock_in
// @Tags Clock_in
// @Summary 批量删除Clock_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Clock_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /clock_in/deleteClock_inByIds [delete]
func (clock_inApi *Clock_inApi) DeleteClock_inByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := clock_inService.DeleteClock_inByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateClock_in 更新Clock_in
// @Tags Clock_in
// @Summary 更新Clock_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Clock_in true "更新Clock_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clock_in/updateClock_in [put]
func (clock_inApi *Clock_inApi) UpdateClock_in(c *gin.Context) {
	var clock_in EpidemicInfo.Clock_in
	_ = c.ShouldBindJSON(&clock_in)
	if err := clock_inService.UpdateClock_in(clock_in); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindClock_in 用id查询Clock_in
// @Tags Clock_in
// @Summary 用id查询Clock_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Clock_in true "用id查询Clock_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clock_in/findClock_in [get]
func (clock_inApi *Clock_inApi) FindClock_in(c *gin.Context) {
	var clock_in EpidemicInfo.Clock_in
	_ = c.ShouldBindQuery(&clock_in)
	if err, reclock_in := clock_inService.GetClock_in(clock_in.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclock_in": reclock_in}, c)
	}
}

// GetClock_inList 分页获取Clock_in列表
// @Tags Clock_in
// @Summary 分页获取Clock_in列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.Clock_inSearch true "分页获取Clock_in列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clock_in/getClock_inList [get]
func (clock_inApi *Clock_inApi) GetClock_inList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.Clock_inSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := clock_inService.GetClock_inInfoList(pageInfo); err != nil {
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
