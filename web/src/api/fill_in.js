import service from '@/utils/request'

// @Tags Fill_in
// @Summary 创建Fill_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Fill_in true "创建Fill_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fill_in/createFill_in [post]
export const createFill_in = (data) => {
  return service({
    url: '/fill_in/createFill_in',
    method: 'post',
    data
  })
}

// @Tags Fill_in
// @Summary 删除Fill_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Fill_in true "删除Fill_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fill_in/deleteFill_in [delete]
export const deleteFill_in = (data) => {
  return service({
    url: '/fill_in/deleteFill_in',
    method: 'delete',
    data
  })
}

// @Tags Fill_in
// @Summary 删除Fill_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Fill_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fill_in/deleteFill_in [delete]
export const deleteFill_inByIds = (data) => {
  return service({
    url: '/fill_in/deleteFill_inByIds',
    method: 'delete',
    data
  })
}

// @Tags Fill_in
// @Summary 更新Fill_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Fill_in true "更新Fill_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fill_in/updateFill_in [put]
export const updateFill_in = (data) => {
  return service({
    url: '/fill_in/updateFill_in',
    method: 'put',
    data
  })
}

// @Tags Fill_in
// @Summary 用id查询Fill_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Fill_in true "用id查询Fill_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fill_in/findFill_in [get]
export const findFill_in = (params) => {
  return service({
    url: '/fill_in/findFill_in',
    method: 'get',
    params
  })
}

// @Tags Fill_in
// @Summary 分页获取Fill_in列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Fill_in列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fill_in/getFill_inList [get]
export const getFill_inList = (params) => {
  return service({
    url: '/fill_in/getFill_inList',
    method: 'get',
    params
  })
}
