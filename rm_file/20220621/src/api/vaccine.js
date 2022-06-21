import service from '@/utils/request'

// @Tags Vaccine
// @Summary 创建Vaccine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vaccine true "创建Vaccine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vaccine/createVaccine [post]
export const createVaccine = (data) => {
  return service({
    url: '/vaccine/createVaccine',
    method: 'post',
    data
  })
}

// @Tags Vaccine
// @Summary 删除Vaccine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vaccine true "删除Vaccine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vaccine/deleteVaccine [delete]
export const deleteVaccine = (data) => {
  return service({
    url: '/vaccine/deleteVaccine',
    method: 'delete',
    data
  })
}

// @Tags Vaccine
// @Summary 删除Vaccine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Vaccine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vaccine/deleteVaccine [delete]
export const deleteVaccineByIds = (data) => {
  return service({
    url: '/vaccine/deleteVaccineByIds',
    method: 'delete',
    data
  })
}

// @Tags Vaccine
// @Summary 更新Vaccine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vaccine true "更新Vaccine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /vaccine/updateVaccine [put]
export const updateVaccine = (data) => {
  return service({
    url: '/vaccine/updateVaccine',
    method: 'put',
    data
  })
}

// @Tags Vaccine
// @Summary 用id查询Vaccine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Vaccine true "用id查询Vaccine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /vaccine/findVaccine [get]
export const findVaccine = (params) => {
  return service({
    url: '/vaccine/findVaccine',
    method: 'get',
    params
  })
}

// @Tags Vaccine
// @Summary 分页获取Vaccine列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Vaccine列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vaccine/getVaccineList [get]
export const getVaccineList = (params) => {
  return service({
    url: '/vaccine/getVaccineList',
    method: 'get',
    params
  })
}
