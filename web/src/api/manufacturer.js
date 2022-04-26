import service from '@/utils/request'

// @Tags Manufacturer
// @Summary 创建Manufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Manufacturer true "创建Manufacturer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /manufacturer/createManufacturer [post]
export const createManufacturer = (data) => {
  return service({
    url: '/manufacturer/createManufacturer',
    method: 'post',
    data
  })
}

// @Tags Manufacturer
// @Summary 删除Manufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Manufacturer true "删除Manufacturer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /manufacturer/deleteManufacturer [delete]
export const deleteManufacturer = (data) => {
  return service({
    url: '/manufacturer/deleteManufacturer',
    method: 'delete',
    data
  })
}

// @Tags Manufacturer
// @Summary 删除Manufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Manufacturer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /manufacturer/deleteManufacturer [delete]
export const deleteManufacturerByIds = (data) => {
  return service({
    url: '/manufacturer/deleteManufacturerByIds',
    method: 'delete',
    data
  })
}

// @Tags Manufacturer
// @Summary 更新Manufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Manufacturer true "更新Manufacturer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /manufacturer/updateManufacturer [put]
export const updateManufacturer = (data) => {
  return service({
    url: '/manufacturer/updateManufacturer',
    method: 'put',
    data
  })
}

// @Tags Manufacturer
// @Summary 用id查询Manufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Manufacturer true "用id查询Manufacturer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /manufacturer/findManufacturer [get]
export const findManufacturer = (params) => {
  return service({
    url: '/manufacturer/findManufacturer',
    method: 'get',
    params
  })
}

// @Tags Manufacturer
// @Summary 分页获取Manufacturer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Manufacturer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /manufacturer/getManufacturerList [get]
export const getManufacturerList = (params) => {
  return service({
    url: '/manufacturer/getManufacturerList',
    method: 'get',
    params
  })
}
