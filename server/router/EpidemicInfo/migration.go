package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MigrationRouter struct {
}

// InitMigrationRouter 初始化 Migration 路由信息
func (s *MigrationRouter) InitMigrationRouter(Router *gin.RouterGroup) {
	migrationRouter := Router.Group("migration").Use(middleware.OperationRecord())
	migrationRouterWithoutRecord := Router.Group("migration")
	var migrationApi = v1.ApiGroupApp.EpidemicInfoApiGroup.MigrationApi
	{
		migrationRouter.POST("createMigration", migrationApi.CreateMigration)   // 新建Migration
		migrationRouter.DELETE("deleteMigration", migrationApi.DeleteMigration) // 删除Migration
		migrationRouter.DELETE("deleteMigrationByIds", migrationApi.DeleteMigrationByIds) // 批量删除Migration
		migrationRouter.PUT("updateMigration", migrationApi.UpdateMigration)    // 更新Migration
	}
	{
		migrationRouterWithoutRecord.GET("findMigration", migrationApi.FindMigration)        // 根据ID获取Migration
		migrationRouterWithoutRecord.GET("getMigrationList", migrationApi.GetMigrationList)  // 获取Migration列表
	}
}
