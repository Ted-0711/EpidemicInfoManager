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

type QuestionnaireApi struct {
}

var questionnaireService = service.ServiceGroupApp.EpidemicInfoServiceGroup.QuestionnaireService


// CreateQuestionnaire 创建Questionnaire
// @Tags Questionnaire
// @Summary 创建Questionnaire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Questionnaire true "创建Questionnaire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionnaire/createQuestionnaire [post]
func (questionnaireApi *QuestionnaireApi) CreateQuestionnaire(c *gin.Context) {
	var questionnaire EpidemicInfo.Questionnaire
	_ = c.ShouldBindJSON(&questionnaire)
	if err := questionnaireService.CreateQuestionnaire(questionnaire); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuestionnaire 删除Questionnaire
// @Tags Questionnaire
// @Summary 删除Questionnaire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Questionnaire true "删除Questionnaire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionnaire/deleteQuestionnaire [delete]
func (questionnaireApi *QuestionnaireApi) DeleteQuestionnaire(c *gin.Context) {
	var questionnaire EpidemicInfo.Questionnaire
	_ = c.ShouldBindJSON(&questionnaire)
	if err := questionnaireService.DeleteQuestionnaire(questionnaire); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuestionnaireByIds 批量删除Questionnaire
// @Tags Questionnaire
// @Summary 批量删除Questionnaire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Questionnaire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /questionnaire/deleteQuestionnaireByIds [delete]
func (questionnaireApi *QuestionnaireApi) DeleteQuestionnaireByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := questionnaireService.DeleteQuestionnaireByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuestionnaire 更新Questionnaire
// @Tags Questionnaire
// @Summary 更新Questionnaire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Questionnaire true "更新Questionnaire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionnaire/updateQuestionnaire [put]
func (questionnaireApi *QuestionnaireApi) UpdateQuestionnaire(c *gin.Context) {
	var questionnaire EpidemicInfo.Questionnaire
	_ = c.ShouldBindJSON(&questionnaire)
	if err := questionnaireService.UpdateQuestionnaire(questionnaire); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuestionnaire 用id查询Questionnaire
// @Tags Questionnaire
// @Summary 用id查询Questionnaire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Questionnaire true "用id查询Questionnaire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionnaire/findQuestionnaire [get]
func (questionnaireApi *QuestionnaireApi) FindQuestionnaire(c *gin.Context) {
	var questionnaire EpidemicInfo.Questionnaire
	_ = c.ShouldBindQuery(&questionnaire)
	if err, requestionnaire := questionnaireService.GetQuestionnaire(questionnaire.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestionnaire": requestionnaire}, c)
	}
}

// GetQuestionnaireList 分页获取Questionnaire列表
// @Tags Questionnaire
// @Summary 分页获取Questionnaire列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.QuestionnaireSearch true "分页获取Questionnaire列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionnaire/getQuestionnaireList [get]
func (questionnaireApi *QuestionnaireApi) GetQuestionnaireList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.QuestionnaireSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := questionnaireService.GetQuestionnaireInfoList(pageInfo); err != nil {
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
