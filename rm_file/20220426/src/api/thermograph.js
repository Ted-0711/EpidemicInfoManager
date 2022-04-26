import service from '@/utils/request'

// @Tags Thermograph
// @Summary 创建Thermograph
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Thermograph true "创建Thermograph"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /thermograph/createThermograph [post]
export const createThermograph = (data) => {
  return service({
    url: '/thermograph/createThermograph',
    method: 'post',
    data
  })
}

// @Tags Thermograph
// @Summary 删除Thermograph
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Thermograph true "删除Thermograph"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /thermograph/deleteThermograph [delete]
export const deleteThermograph = (data) => {
  return service({
    url: '/thermograph/deleteThermograph',
    method: 'delete',
    data
  })
}

// @Tags Thermograph
// @Summary 删除Thermograph
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Thermograph"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /thermograph/deleteThermograph [delete]
export const deleteThermographByIds = (data) => {
  return service({
    url: '/thermograph/deleteThermographByIds',
    method: 'delete',
    data
  })
}

// @Tags Thermograph
// @Summary 更新Thermograph
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Thermograph true "更新Thermograph"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /thermograph/updateThermograph [put]
export const updateThermograph = (data) => {
  return service({
    url: '/thermograph/updateThermograph',
    method: 'put',
    data
  })
}

// @Tags Thermograph
// @Summary 用id查询Thermograph
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Thermograph true "用id查询Thermograph"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /thermograph/findThermograph [get]
export const findThermograph = (params) => {
  return service({
    url: '/thermograph/findThermograph',
    method: 'get',
    params
  })
}

// @Tags Thermograph
// @Summary 分页获取Thermograph列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Thermograph列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /thermograph/getThermographList [get]
export const getThermographList = (params) => {
  return service({
    url: '/thermograph/getThermographList',
    method: 'get',
    params
  })
}
