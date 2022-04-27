import service from '@/utils/request'

// @Tags Counselor
// @Summary 创建Counselor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Counselor true "创建Counselor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /counselor/createCounselor [post]
export const createCounselor = (data) => {
  return service({
    url: '/counselor/createCounselor',
    method: 'post',
    data
  })
}

// @Tags Counselor
// @Summary 删除Counselor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Counselor true "删除Counselor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /counselor/deleteCounselor [delete]
export const deleteCounselor = (data) => {
  return service({
    url: '/counselor/deleteCounselor',
    method: 'delete',
    data
  })
}

// @Tags Counselor
// @Summary 删除Counselor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Counselor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /counselor/deleteCounselor [delete]
export const deleteCounselorByIds = (data) => {
  return service({
    url: '/counselor/deleteCounselorByIds',
    method: 'delete',
    data
  })
}

// @Tags Counselor
// @Summary 更新Counselor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Counselor true "更新Counselor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /counselor/updateCounselor [put]
export const updateCounselor = (data) => {
  return service({
    url: '/counselor/updateCounselor',
    method: 'put',
    data
  })
}

// @Tags Counselor
// @Summary 用id查询Counselor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Counselor true "用id查询Counselor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /counselor/findCounselor [get]
export const findCounselor = (params) => {
  return service({
    url: '/counselor/findCounselor',
    method: 'get',
    params
  })
}

// @Tags Counselor
// @Summary 分页获取Counselor列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Counselor列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /counselor/getCounselorList [get]
export const getCounselorList = (params) => {
  return service({
    url: '/counselor/getCounselorList',
    method: 'get',
    params
  })
}
