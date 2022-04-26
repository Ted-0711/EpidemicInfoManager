package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type QuarantineService struct {
}

// CreateQuarantine 创建Quarantine记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantineService *QuarantineService) CreateQuarantine(quarantine EpidemicInfo.Quarantine) (err error) {
	err = global.GVA_DB.Create(&quarantine).Error
	return err
}

// DeleteQuarantine 删除Quarantine记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantineService *QuarantineService) DeleteQuarantine(quarantine EpidemicInfo.Quarantine) (err error) {
	err = global.GVA_DB.Delete(&quarantine).Error
	return err
}

// DeleteQuarantineByIds 批量删除Quarantine记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantineService *QuarantineService) DeleteQuarantineByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Quarantine{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateQuarantine 更新Quarantine记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantineService *QuarantineService) UpdateQuarantine(quarantine EpidemicInfo.Quarantine) (err error) {
	err = global.GVA_DB.Save(&quarantine).Error
	return err
}

// GetQuarantine 根据id获取Quarantine记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantineService *QuarantineService) GetQuarantine(id uint) (err error, quarantine EpidemicInfo.Quarantine) {
	err = global.GVA_DB.Where("id = ?", id).First(&quarantine).Error
	return
}

// GetQuarantineInfoList 分页获取Quarantine记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantineService *QuarantineService) GetQuarantineInfoList(info EpidemicInfoReq.QuarantineSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Quarantine{})
	var quarantines []EpidemicInfo.Quarantine
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Student_id != "" {
		db = db.Where("student_id LIKE ?", "%"+info.Student_id+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&quarantines).Error
	return err, quarantines, total
}
