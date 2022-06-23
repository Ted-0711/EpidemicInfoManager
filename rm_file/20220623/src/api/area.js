import service from '@/utils/request'

// @Tags Area
// @Summary 创建Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Area true "创建Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /area/createArea [post]
export const createArea = (data) => {
  return service({
    url: '/area/createArea',
    method: 'post',
    data
  })
}

// @Tags Area
// @Summary 删除Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Area true "删除Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /area/deleteArea [delete]
export const deleteArea = (data) => {
  return service({
    url: '/area/deleteArea',
    method: 'delete',
    data
  })
}

// @Tags Area
// @Summary 删除Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /area/deleteArea [delete]
export const deleteAreaByIds = (data) => {
  return service({
    url: '/area/deleteAreaByIds',
    method: 'delete',
    data
  })
}

// @Tags Area
// @Summary 更新Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Area true "更新Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /area/updateArea [put]
export const updateArea = (data) => {
  return service({
    url: '/area/updateArea',
    method: 'put',
    data
  })
}

// @Tags Area
// @Summary 用id查询Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Area true "用id查询Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /area/findArea [get]
export const findArea = (params) => {
  return service({
    url: '/area/findArea',
    method: 'get',
    params
  })
}

// @Tags Area
// @Summary 分页获取Area列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Area列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /area/getAreaList [get]
export const getAreaList = (params) => {
  return service({
    url: '/area/getAreaList',
    method: 'get',
    params
  })
}
