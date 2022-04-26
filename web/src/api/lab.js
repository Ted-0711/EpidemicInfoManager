import service from '@/utils/request'

// @Tags Lab
// @Summary 创建Lab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lab true "创建Lab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lab/createLab [post]
export const createLab = (data) => {
  return service({
    url: '/lab/createLab',
    method: 'post',
    data
  })
}

// @Tags Lab
// @Summary 删除Lab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lab true "删除Lab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lab/deleteLab [delete]
export const deleteLab = (data) => {
  return service({
    url: '/lab/deleteLab',
    method: 'delete',
    data
  })
}

// @Tags Lab
// @Summary 删除Lab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Lab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lab/deleteLab [delete]
export const deleteLabByIds = (data) => {
  return service({
    url: '/lab/deleteLabByIds',
    method: 'delete',
    data
  })
}

// @Tags Lab
// @Summary 更新Lab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Lab true "更新Lab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lab/updateLab [put]
export const updateLab = (data) => {
  return service({
    url: '/lab/updateLab',
    method: 'put',
    data
  })
}

// @Tags Lab
// @Summary 用id查询Lab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Lab true "用id查询Lab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lab/findLab [get]
export const findLab = (params) => {
  return service({
    url: '/lab/findLab',
    method: 'get',
    params
  })
}

// @Tags Lab
// @Summary 分页获取Lab列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Lab列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lab/getLabList [get]
export const getLabList = (params) => {
  return service({
    url: '/lab/getLabList',
    method: 'get',
    params
  })
}
