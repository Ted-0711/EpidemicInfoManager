package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type MeasureService struct {
}

// CreateMeasure 创建Measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (measureService *MeasureService) CreateMeasure(measure EpidemicInfo.Measure) (err error) {
	err = global.GVA_DB.Create(&measure).Error
	return err
}

// DeleteMeasure 删除Measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (measureService *MeasureService)DeleteMeasure(measure EpidemicInfo.Measure) (err error) {
	err = global.GVA_DB.Delete(&measure).Error
	return err
}

// DeleteMeasureByIds 批量删除Measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (measureService *MeasureService)DeleteMeasureByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Measure{},"id in ?",ids.Ids).Error
	return err
}

// UpdateMeasure 更新Measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (measureService *MeasureService)UpdateMeasure(measure EpidemicInfo.Measure) (err error) {
	err = global.GVA_DB.Save(&measure).Error
	return err
}

// GetMeasure 根据id获取Measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (measureService *MeasureService)GetMeasure(id uint) (err error, measure EpidemicInfo.Measure) {
	err = global.GVA_DB.Where("id = ?", id).First(&measure).Error
	return
}

// GetMeasureInfoList 分页获取Measure记录
// Author [piexlmax](https://github.com/piexlmax)
func (measureService *MeasureService)GetMeasureInfoList(info EpidemicInfoReq.MeasureSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Measure{})
    var measures []EpidemicInfo.Measure
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&measures).Error
	return err, measures, total
}
