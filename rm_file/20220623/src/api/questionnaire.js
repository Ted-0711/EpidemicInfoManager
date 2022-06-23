import service from '@/utils/request'

// @Tags Questionnaire
// @Summary 创建Questionnaire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Questionnaire true "创建Questionnaire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionnaire/createQuestionnaire [post]
export const createQuestionnaire = (data) => {
  return service({
    url: '/questionnaire/createQuestionnaire',
    method: 'post',
    data
  })
}

// @Tags Questionnaire
// @Summary 删除Questionnaire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Questionnaire true "删除Questionnaire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionnaire/deleteQuestionnaire [delete]
export const deleteQuestionnaire = (data) => {
  return service({
    url: '/questionnaire/deleteQuestionnaire',
    method: 'delete',
    data
  })
}

// @Tags Questionnaire
// @Summary 删除Questionnaire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Questionnaire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionnaire/deleteQuestionnaire [delete]
export const deleteQuestionnaireByIds = (data) => {
  return service({
    url: '/questionnaire/deleteQuestionnaireByIds',
    method: 'delete',
    data
  })
}

// @Tags Questionnaire
// @Summary 更新Questionnaire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Questionnaire true "更新Questionnaire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionnaire/updateQuestionnaire [put]
export const updateQuestionnaire = (data) => {
  return service({
    url: '/questionnaire/updateQuestionnaire',
    method: 'put',
    data
  })
}

// @Tags Questionnaire
// @Summary 用id查询Questionnaire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Questionnaire true "用id查询Questionnaire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionnaire/findQuestionnaire [get]
export const findQuestionnaire = (params) => {
  return service({
    url: '/questionnaire/findQuestionnaire',
    method: 'get',
    params
  })
}

// @Tags Questionnaire
// @Summary 分页获取Questionnaire列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Questionnaire列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionnaire/getQuestionnaireList [get]
export const getQuestionnaireList = (params) => {
  return service({
    url: '/questionnaire/getQuestionnaireList',
    method: 'get',
    params
  })
}
