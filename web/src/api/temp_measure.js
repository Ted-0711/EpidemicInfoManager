import service from '@/utils/request'

// @Tags Temp_measure
// @Summary 创建Temp_measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Temp_measure true "创建Temp_measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /temp_measure/createTemp_measure [post]
export const createTemp_measure = (data) => {
  return service({
    url: '/temp_measure/createTemp_measure',
    method: 'post',
    data
  })
}

// @Tags Temp_measure
// @Summary 删除Temp_measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Temp_measure true "删除Temp_measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /temp_measure/deleteTemp_measure [delete]
export const deleteTemp_measure = (data) => {
  return service({
    url: '/temp_measure/deleteTemp_measure',
    method: 'delete',
    data
  })
}

// @Tags Temp_measure
// @Summary 删除Temp_measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Temp_measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /temp_measure/deleteTemp_measure [delete]
export const deleteTemp_measureByIds = (data) => {
  return service({
    url: '/temp_measure/deleteTemp_measureByIds',
    method: 'delete',
    data
  })
}

// @Tags Temp_measure
// @Summary 更新Temp_measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Temp_measure true "更新Temp_measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /temp_measure/updateTemp_measure [put]
export const updateTemp_measure = (data) => {
  return service({
    url: '/temp_measure/updateTemp_measure',
    method: 'put',
    data
  })
}

// @Tags Temp_measure
// @Summary 用id查询Temp_measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Temp_measure true "用id查询Temp_measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /temp_measure/findTemp_measure [get]
export const findTemp_measure = (params) => {
  return service({
    url: '/temp_measure/findTemp_measure',
    method: 'get',
    params
  })
}

// @Tags Temp_measure
// @Summary 分页获取Temp_measure列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Temp_measure列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /temp_measure/getTemp_measureList [get]
export const getTemp_measureList = (params) => {
  return service({
    url: '/temp_measure/getTemp_measureList',
    method: 'get',
    params
  })
}
