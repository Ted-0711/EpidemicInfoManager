import service from '@/utils/request'

// @Tags Notice
// @Summary 创建Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notice true "创建Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notice/createNotice [post]
export const createNotice = (data) => {
  return service({
    url: '/notice/createNotice',
    method: 'post',
    data
  })
}

// @Tags Notice
// @Summary 删除Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notice true "删除Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notice/deleteNotice [delete]
export const deleteNotice = (data) => {
  return service({
    url: '/notice/deleteNotice',
    method: 'delete',
    data
  })
}

// @Tags Notice
// @Summary 删除Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notice/deleteNotice [delete]
export const deleteNoticeByIds = (data) => {
  return service({
    url: '/notice/deleteNoticeByIds',
    method: 'delete',
    data
  })
}

// @Tags Notice
// @Summary 更新Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notice true "更新Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /notice/updateNotice [put]
export const updateNotice = (data) => {
  return service({
    url: '/notice/updateNotice',
    method: 'put',
    data
  })
}

// @Tags Notice
// @Summary 用id查询Notice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Notice true "用id查询Notice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /notice/findNotice [get]
export const findNotice = (params) => {
  return service({
    url: '/notice/findNotice',
    method: 'get',
    params
  })
}

// @Tags Notice
// @Summary 分页获取Notice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Notice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notice/getNoticeList [get]
export const getNoticeList = (params) => {
  return service({
    url: '/notice/getNoticeList',
    method: 'get',
    params
  })
}
