import service from '@/utils/request'

// @Tags Clock_in
// @Summary 创建Clock_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Clock_in true "创建Clock_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clock_in/createClock_in [post]
export const createClock_in = (data) => {
  return service({
    url: '/clock_in/createClock_in',
    method: 'post',
    data
  })
}

// @Tags Clock_in
// @Summary 删除Clock_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Clock_in true "删除Clock_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clock_in/deleteClock_in [delete]
export const deleteClock_in = (data) => {
  return service({
    url: '/clock_in/deleteClock_in',
    method: 'delete',
    data
  })
}

// @Tags Clock_in
// @Summary 删除Clock_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Clock_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clock_in/deleteClock_in [delete]
export const deleteClock_inByIds = (data) => {
  return service({
    url: '/clock_in/deleteClock_inByIds',
    method: 'delete',
    data
  })
}

// @Tags Clock_in
// @Summary 更新Clock_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Clock_in true "更新Clock_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clock_in/updateClock_in [put]
export const updateClock_in = (data) => {
  return service({
    url: '/clock_in/updateClock_in',
    method: 'put',
    data
  })
}

// @Tags Clock_in
// @Summary 用id查询Clock_in
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Clock_in true "用id查询Clock_in"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clock_in/findClock_in [get]
export const findClock_in = (params) => {
  return service({
    url: '/clock_in/findClock_in',
    method: 'get',
    params
  })
}

// @Tags Clock_in
// @Summary 分页获取Clock_in列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Clock_in列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clock_in/getClock_inList [get]
export const getClock_inList = (params) => {
  return service({
    url: '/clock_in/getClock_inList',
    method: 'get',
    params
  })
}
