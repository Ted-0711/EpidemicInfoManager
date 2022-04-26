import service from '@/utils/request'

// @Tags Quarantine
// @Summary 创建Quarantine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Quarantine true "创建Quarantine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /quarantine/createQuarantine [post]
export const createQuarantine = (data) => {
  return service({
    url: '/quarantine/createQuarantine',
    method: 'post',
    data
  })
}

// @Tags Quarantine
// @Summary 删除Quarantine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Quarantine true "删除Quarantine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /quarantine/deleteQuarantine [delete]
export const deleteQuarantine = (data) => {
  return service({
    url: '/quarantine/deleteQuarantine',
    method: 'delete',
    data
  })
}

// @Tags Quarantine
// @Summary 删除Quarantine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Quarantine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /quarantine/deleteQuarantine [delete]
export const deleteQuarantineByIds = (data) => {
  return service({
    url: '/quarantine/deleteQuarantineByIds',
    method: 'delete',
    data
  })
}

// @Tags Quarantine
// @Summary 更新Quarantine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Quarantine true "更新Quarantine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /quarantine/updateQuarantine [put]
export const updateQuarantine = (data) => {
  return service({
    url: '/quarantine/updateQuarantine',
    method: 'put',
    data
  })
}

// @Tags Quarantine
// @Summary 用id查询Quarantine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Quarantine true "用id查询Quarantine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /quarantine/findQuarantine [get]
export const findQuarantine = (params) => {
  return service({
    url: '/quarantine/findQuarantine',
    method: 'get',
    params
  })
}

// @Tags Quarantine
// @Summary 分页获取Quarantine列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Quarantine列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /quarantine/getQuarantineList [get]
export const getQuarantineList = (params) => {
  return service({
    url: '/quarantine/getQuarantineList',
    method: 'get',
    params
  })
}
