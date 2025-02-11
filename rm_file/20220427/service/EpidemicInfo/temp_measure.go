package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type Temp_measureService struct {
}

// CreateTemp_measure 创建Temp_measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (temp_measureService *Temp_measureService) CreateTemp_measure(temp_measure EpidemicInfo.Temp_measure) (err error) {
	err = global.GVA_DB.Create(&temp_measure).Error
	return err
}

// DeleteTemp_measure 删除Temp_measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (temp_measureService *Temp_measureService) DeleteTemp_measure(temp_measure EpidemicInfo.Temp_measure) (err error) {
	err = global.GVA_DB.Delete(&temp_measure).Error
	return err
}

// DeleteTemp_measureByIds 批量删除Temp_measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (temp_measureService *Temp_measureService) DeleteTemp_measureByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Temp_measure{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTemp_measure 更新Temp_measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (temp_measureService *Temp_measureService) UpdateTemp_measure(temp_measure EpidemicInfo.Temp_measure) (err error) {
	err = global.GVA_DB.Save(&temp_measure).Error
	return err
}

// GetTemp_measure 根据id获取Temp_measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (temp_measureService *Temp_measureService) GetTemp_measure(id uint) (err error, temp_measure EpidemicInfo.Temp_measure) {
	err = global.GVA_DB.Where("id = ?", id).First(&temp_measure).Error
	return
}

// GetTemp_measureInfoList 分页获取Temp_measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (temp_measureService *Temp_measureService) GetTemp_measureInfoList(info EpidemicInfoReq.Temp_measureSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Temp_measure{})
	var temp_measures []EpidemicInfo.Temp_measure
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Therm_id != nil {
		db = db.Where("therm_id = ?", info.Therm_id)
	}
	if info.Temperature != nil {
		db = db.Where("temperature > ?", info.Temperature)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&temp_measures).Error
	return err, temp_measures, total
}
