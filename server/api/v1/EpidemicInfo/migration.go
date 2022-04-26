package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type MigrationApi struct {
}

var migrationService = service.ServiceGroupApp.EpidemicInfoServiceGroup.MigrationService


// CreateMigration 创建Migration
// @Tags Migration
// @Summary 创建Migration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Migration true "创建Migration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /migration/createMigration [post]
func (migrationApi *MigrationApi) CreateMigration(c *gin.Context) {
	var migration EpidemicInfo.Migration
	_ = c.ShouldBindJSON(&migration)
	if err := migrationService.CreateMigration(migration); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMigration 删除Migration
// @Tags Migration
// @Summary 删除Migration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Migration true "删除Migration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /migration/deleteMigration [delete]
func (migrationApi *MigrationApi) DeleteMigration(c *gin.Context) {
	var migration EpidemicInfo.Migration
	_ = c.ShouldBindJSON(&migration)
	if err := migrationService.DeleteMigration(migration); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMigrationByIds 批量删除Migration
// @Tags Migration
// @Summary 批量删除Migration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Migration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /migration/deleteMigrationByIds [delete]
func (migrationApi *MigrationApi) DeleteMigrationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := migrationService.DeleteMigrationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMigration 更新Migration
// @Tags Migration
// @Summary 更新Migration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EpidemicInfo.Migration true "更新Migration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /migration/updateMigration [put]
func (migrationApi *MigrationApi) UpdateMigration(c *gin.Context) {
	var migration EpidemicInfo.Migration
	_ = c.ShouldBindJSON(&migration)
	if err := migrationService.UpdateMigration(migration); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMigration 用id查询Migration
// @Tags Migration
// @Summary 用id查询Migration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfo.Migration true "用id查询Migration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /migration/findMigration [get]
func (migrationApi *MigrationApi) FindMigration(c *gin.Context) {
	var migration EpidemicInfo.Migration
	_ = c.ShouldBindQuery(&migration)
	if err, remigration := migrationService.GetMigration(migration.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remigration": remigration}, c)
	}
}

// GetMigrationList 分页获取Migration列表
// @Tags Migration
// @Summary 分页获取Migration列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EpidemicInfoReq.MigrationSearch true "分页获取Migration列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /migration/getMigrationList [get]
func (migrationApi *MigrationApi) GetMigrationList(c *gin.Context) {
	var pageInfo EpidemicInfoReq.MigrationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := migrationService.GetMigrationInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
