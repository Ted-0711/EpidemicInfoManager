package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type ManufacturerService struct {
}

// CreateManufacturer 创建Manufacturer记录
// Author [piexlmax](https://github.com/piexlmax)
func (manufacturerService *ManufacturerService) CreateManufacturer(manufacturer EpidemicInfo.Manufacturer) (err error) {
	err = global.GVA_DB.Create(&manufacturer).Error
	return err
}

// DeleteManufacturer 删除Manufacturer记录
// Author [piexlmax](https://github.com/piexlmax)
func (manufacturerService *ManufacturerService)DeleteManufacturer(manufacturer EpidemicInfo.Manufacturer) (err error) {
	err = global.GVA_DB.Delete(&manufacturer).Error
	return err
}

// DeleteManufacturerByIds 批量删除Manufacturer记录
// Author [piexlmax](https://github.com/piexlmax)
func (manufacturerService *ManufacturerService)DeleteManufacturerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Manufacturer{},"id in ?",ids.Ids).Error
	return err
}

// UpdateManufacturer 更新Manufacturer记录
// Author [piexlmax](https://github.com/piexlmax)
func (manufacturerService *ManufacturerService)UpdateManufacturer(manufacturer EpidemicInfo.Manufacturer) (err error) {
	err = global.GVA_DB.Save(&manufacturer).Error
	return err
}

// GetManufacturer 根据id获取Manufacturer记录
// Author [piexlmax](https://github.com/piexlmax)
func (manufacturerService *ManufacturerService)GetManufacturer(id uint) (err error, manufacturer EpidemicInfo.Manufacturer) {
	err = global.GVA_DB.Where("id = ?", id).First(&manufacturer).Error
	return
}

// GetManufacturerInfoList 分页获取Manufacturer记录
// Author [piexlmax](https://github.com/piexlmax)
func (manufacturerService *ManufacturerService)GetManufacturerInfoList(info EpidemicInfoReq.ManufacturerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Manufacturer{})
    var manufacturers []EpidemicInfo.Manufacturer
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Area_id != nil {
        db = db.Where("area_id = ?",info.Area_id)
    }
    if info.Mfr_name != "" {
        db = db.Where("mfr_name LIKE ?","%"+ info.Mfr_name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&manufacturers).Error
	return err, manufacturers, total
}
