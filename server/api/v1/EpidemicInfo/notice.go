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

type NoticeApi struct {
}

var noticeService = service.ServiceGroupApp.EpidemicInfoServiceGroup.NoticeService


// CreateNotice 创建Notice
// @Tags Notice
// @Summary 创建Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Notice true "创建Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notice/createNotice [post]
func (noticeApi *NoticeApi) CreateNotice(c *gin.Context) {
	var notice EpidemicInfo.Notice
	_ = c.ShouldBindJSON(&notice)
	if err := noticeService.CreateNotice(notice); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteNotice 删除Notice
// @Tags Notice
// @Summary 删除Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Notice true "删除Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notice/deleteNotice [delete]
func (noticeApi *NoticeApi) DeleteNotice(c *gin.Context) {
	var notice EpidemicInfo.Notice
	_ = c.ShouldBindJSON(&notice)
	if err := noticeService.DeleteNotice(notice); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNoticeByIds 批量删除Notice
// @Tags Notice
// @Summary 批量删除Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /notice/deleteNoticeByIds [delete]
func (noticeApi *NoticeApi) DeleteNoticeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := noticeService.DeleteNoticeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNotice 更新Notice
// @Tags Notice
// @Summary 更新Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Notice true "更新Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /notice/updateNotice [put]
func (noticeApi *NoticeApi) UpdateNotice(c *gin.Context) {
	var notice EpidemicInfo.Notice
	_ = c.ShouldBindJSON(&notice)
	if err := noticeService.UpdateNotice(notice); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNotice 用id查询Notice
// @Tags Notice
// @Summary 用id查询Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Notice true "用id查询Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /notice/findNotice [get]
func (noticeApi *NoticeApi) FindNotice(c *gin.Context) {
	var notice EpidemicInfo.Notice
	_ = c.ShouldBindQuery(&notice)
	if err, renotice := noticeService.GetNotice(notice.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"renotice": renotice}, c)
	}
}

// GetNoticeList 分页获取Notice列表
// @Tags Notice
// @Summary 分页获取Notice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.NoticeSearch true "分页获取Notice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notice/getNoticeList [get]
func (noticeApi *NoticeApi) GetNoticeList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.NoticeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := noticeService.GetNoticeInfoList(pageInfo); err != nil {
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
