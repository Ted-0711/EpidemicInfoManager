import service from '@/utils/request'

// @Tags Test_swab
// @Summary 创建Test_swab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test_swab true "创建Test_swab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test_swab/createTest_swab [post]
export const createTest_swab = (data) => {
  return service({
    url: '/test_swab/createTest_swab',
    method: 'post',
    data
  })
}

// @Tags Test_swab
// @Summary 删除Test_swab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test_swab true "删除Test_swab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test_swab/deleteTest_swab [delete]
export const deleteTest_swab = (data) => {
  return service({
    url: '/test_swab/deleteTest_swab',
    method: 'delete',
    data
  })
}

// @Tags Test_swab
// @Summary 删除Test_swab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Test_swab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test_swab/deleteTest_swab [delete]
export const deleteTest_swabByIds = (data) => {
  return service({
    url: '/test_swab/deleteTest_swabByIds',
    method: 'delete',
    data
  })
}

// @Tags Test_swab
// @Summary 更新Test_swab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test_swab true "更新Test_swab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /test_swab/updateTest_swab [put]
export const updateTest_swab = (data) => {
  return service({
    url: '/test_swab/updateTest_swab',
    method: 'put',
    data
  })
}

// @Tags Test_swab
// @Summary 用id查询Test_swab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Test_swab true "用id查询Test_swab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /test_swab/findTest_swab [get]
export const findTest_swab = (params) => {
  return service({
    url: '/test_swab/findTest_swab',
    method: 'get',
    params
  })
}

// @Tags Test_swab
// @Summary 分页获取Test_swab列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Test_swab列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test_swab/getTest_swabList [get]
export const getTest_swabList = (params) => {
  return service({
    url: '/test_swab/getTest_swabList',
    method: 'get',
    params
  })
}
