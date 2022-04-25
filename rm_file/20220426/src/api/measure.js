import service from '@/utils/request'

// @Tags Measure
// @Summary 创建Measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Measure true "创建Measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /measure/createMeasure [post]
export const createMeasure = (data) => {
  return service({
    url: '/measure/createMeasure',
    method: 'post',
    data
  })
}

// @Tags Measure
// @Summary 删除Measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Measure true "删除Measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /measure/deleteMeasure [delete]
export const deleteMeasure = (data) => {
  return service({
    url: '/measure/deleteMeasure',
    method: 'delete',
    data
  })
}

// @Tags Measure
// @Summary 删除Measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /measure/deleteMeasure [delete]
export const deleteMeasureByIds = (data) => {
  return service({
    url: '/measure/deleteMeasureByIds',
    method: 'delete',
    data
  })
}

// @Tags Measure
// @Summary 更新Measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Measure true "更新Measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /measure/updateMeasure [put]
export const updateMeasure = (data) => {
  return service({
    url: '/measure/updateMeasure',
    method: 'put',
    data
  })
}

// @Tags Measure
// @Summary 用id查询Measure
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Measure true "用id查询Measure"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /measure/findMeasure [get]
export const findMeasure = (params) => {
  return service({
    url: '/measure/findMeasure',
    method: 'get',
    params
  })
}

// @Tags Measure
// @Summary 分页获取Measure列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Measure列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /measure/getMeasureList [get]
export const getMeasureList = (params) => {
  return service({
    url: '/measure/getMeasureList',
    method: 'get',
    params
  })
}
