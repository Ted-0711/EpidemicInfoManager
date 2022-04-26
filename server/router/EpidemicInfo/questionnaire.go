package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QuestionnaireRouter struct {
}

// InitQuestionnaireRouter 初始化 Questionnaire 路由信息
func (s *QuestionnaireRouter) InitQuestionnaireRouter(Router *gin.RouterGroup) {
	questionnaireRouter := Router.Group("questionnaire").Use(middleware.OperationRecord())
	questionnaireRouterWithoutRecord := Router.Group("questionnaire")
	var questionnaireApi = v1.ApiGroupApp.EpidemicInfoApiGroup.QuestionnaireApi
	{
		questionnaireRouter.POST("createQuestionnaire", questionnaireApi.CreateQuestionnaire)   // 新建Questionnaire
		questionnaireRouter.DELETE("deleteQuestionnaire", questionnaireApi.DeleteQuestionnaire) // 删除Questionnaire
		questionnaireRouter.DELETE("deleteQuestionnaireByIds", questionnaireApi.DeleteQuestionnaireByIds) // 批量删除Questionnaire
		questionnaireRouter.PUT("updateQuestionnaire", questionnaireApi.UpdateQuestionnaire)    // 更新Questionnaire
	}
	{
		questionnaireRouterWithoutRecord.GET("findQuestionnaire", questionnaireApi.FindQuestionnaire)        // 根据ID获取Questionnaire
		questionnaireRouterWithoutRecord.GET("getQuestionnaireList", questionnaireApi.GetQuestionnaireList)  // 获取Questionnaire列表
	}
}
