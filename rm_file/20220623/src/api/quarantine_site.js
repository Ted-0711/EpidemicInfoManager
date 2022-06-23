import service from '@/utils/request'

// @Tags Quarantine_site
// @Summary 创建Quarantine_site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Quarantine_site true "创建Quarantine_site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /quarantine_site/createQuarantine_site [post]
export const createQuarantine_site = (data) => {
  return service({
    url: '/quarantine_site/createQuarantine_site',
    method: 'post',
    data
  })
}

// @Tags Quarantine_site
// @Summary 删除Quarantine_site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Quarantine_site true "删除Quarantine_site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /quarantine_site/deleteQuarantine_site [delete]
export const deleteQuarantine_site = (data) => {
  return service({
    url: '/quarantine_site/deleteQuarantine_site',
    method: 'delete',
    data
  })
}

// @Tags Quarantine_site
// @Summary 删除Quarantine_site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Quarantine_site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /quarantine_site/deleteQuarantine_site [delete]
export const deleteQuarantine_siteByIds = (data) => {
  return service({
    url: '/quarantine_site/deleteQuarantine_siteByIds',
    method: 'delete',
    data
  })
}

// @Tags Quarantine_site
// @Summary 更新Quarantine_site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Quarantine_site true "更新Quarantine_site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /quarantine_site/updateQuarantine_site [put]
export const updateQuarantine_site = (data) => {
  return service({
    url: '/quarantine_site/updateQuarantine_site',
    method: 'put',
    data
  })
}

// @Tags Quarantine_site
// @Summary 用id查询Quarantine_site
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Quarantine_site true "用id查询Quarantine_site"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /quarantine_site/findQuarantine_site [get]
export const findQuarantine_site = (params) => {
  return service({
    url: '/quarantine_site/findQuarantine_site',
    method: 'get',
    params
  })
}

// @Tags Quarantine_site
// @Summary 分页获取Quarantine_site列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Quarantine_site列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /quarantine_site/getQuarantine_siteList [get]
export const getQuarantine_siteList = (params) => {
  return service({
    url: '/quarantine_site/getQuarantine_siteList',
    method: 'get',
    params
  })
}
