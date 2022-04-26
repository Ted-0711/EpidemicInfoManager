import service from '@/utils/request'

// @Tags Migration
// @Summary 创建Migration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Migration true "创建Migration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /migration/createMigration [post]
export const createMigration = (data) => {
  return service({
    url: '/migration/createMigration',
    method: 'post',
    data
  })
}

// @Tags Migration
// @Summary 删除Migration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Migration true "删除Migration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /migration/deleteMigration [delete]
export const deleteMigration = (data) => {
  return service({
    url: '/migration/deleteMigration',
    method: 'delete',
    data
  })
}

// @Tags Migration
// @Summary 删除Migration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Migration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /migration/deleteMigration [delete]
export const deleteMigrationByIds = (data) => {
  return service({
    url: '/migration/deleteMigrationByIds',
    method: 'delete',
    data
  })
}

// @Tags Migration
// @Summary 更新Migration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Migration true "更新Migration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /migration/updateMigration [put]
export const updateMigration = (data) => {
  return service({
    url: '/migration/updateMigration',
    method: 'put',
    data
  })
}

// @Tags Migration
// @Summary 用id查询Migration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Migration true "用id查询Migration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /migration/findMigration [get]
export const findMigration = (params) => {
  return service({
    url: '/migration/findMigration',
    method: 'get',
    params
  })
}

// @Tags Migration
// @Summary 分页获取Migration列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Migration列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /migration/getMigrationList [get]
export const getMigrationList = (params) => {
  return service({
    url: '/migration/getMigrationList',
    method: 'get',
    params
  })
}
