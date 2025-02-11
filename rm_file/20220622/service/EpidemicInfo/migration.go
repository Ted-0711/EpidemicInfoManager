package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type MigrationService struct {
}

// CreateMigration 创建Migration记录
// Author [piexlmax](https://github.com/piexlmax)
func (migrationService *MigrationService) CreateMigration(migration EpidemicInfo.Migration) (err error) {
	err = global.GVA_DB.Create(&migration).Error
	return err
}

// DeleteMigration 删除Migration记录
// Author [piexlmax](https://github.com/piexlmax)
func (migrationService *MigrationService)DeleteMigration(migration EpidemicInfo.Migration) (err error) {
	err = global.GVA_DB.Delete(&migration).Error
	return err
}

// DeleteMigrationByIds 批量删除Migration记录
// Author [piexlmax](https://github.com/piexlmax)
func (migrationService *MigrationService)DeleteMigrationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Migration{},"id in ?",ids.Ids).Error
	return err
}

// UpdateMigration 更新Migration记录
// Author [piexlmax](https://github.com/piexlmax)
func (migrationService *MigrationService)UpdateMigration(migration EpidemicInfo.Migration) (err error) {
	err = global.GVA_DB.Save(&migration).Error
	return err
}

// GetMigration 根据id获取Migration记录
// Author [piexlmax](https://github.com/piexlmax)
func (migrationService *MigrationService)GetMigration(id uint) (err error, migration EpidemicInfo.Migration) {
	err = global.GVA_DB.Where("id = ?", id).First(&migration).Error
	return
}

// GetMigrationInfoList 分页获取Migration记录
// Author [piexlmax](https://github.com/piexlmax)
func (migrationService *MigrationService)GetMigrationInfoList(info EpidemicInfoReq.MigrationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Migration{})
    var migrations []EpidemicInfo.Migration
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Student_id != "" {
        db = db.Where("student_id LIKE ?","%"+ info.Student_id+"%")
    }
    if info.Start_area_id != nil {
        db = db.Where("start_area_id = ?",info.Start_area_id)
    }
    if info.Des_area_id != nil {
        db = db.Where("des_area_id = ?",info.Des_area_id)
    }
    if info.Vehicle_info != "" {
        db = db.Where("vehicle_info LIKE ?","%"+ info.Vehicle_info+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&migrations).Error
	return err, migrations, total
}
