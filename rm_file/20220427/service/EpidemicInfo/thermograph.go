package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type ThermographService struct {
}

// CreateThermograph 创建Thermograph记录
// Author [piexlmax](https://github.com/piexlmax)
func (thermographService *ThermographService) CreateThermograph(thermograph EpidemicInfo.Thermograph) (err error) {
	err = global.GVA_DB.Create(&thermograph).Error
	return err
}

// DeleteThermograph 删除Thermograph记录
// Author [piexlmax](https://github.com/piexlmax)
func (thermographService *ThermographService)DeleteThermograph(thermograph EpidemicInfo.Thermograph) (err error) {
	err = global.GVA_DB.Delete(&thermograph).Error
	return err
}

// DeleteThermographByIds 批量删除Thermograph记录
// Author [piexlmax](https://github.com/piexlmax)
func (thermographService *ThermographService)DeleteThermographByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Thermograph{},"id in ?",ids.Ids).Error
	return err
}

// UpdateThermograph 更新Thermograph记录
// Author [piexlmax](https://github.com/piexlmax)
func (thermographService *ThermographService)UpdateThermograph(thermograph EpidemicInfo.Thermograph) (err error) {
	err = global.GVA_DB.Save(&thermograph).Error
	return err
}

// GetThermograph 根据id获取Thermograph记录
// Author [piexlmax](https://github.com/piexlmax)
func (thermographService *ThermographService)GetThermograph(id uint) (err error, thermograph EpidemicInfo.Thermograph) {
	err = global.GVA_DB.Where("id = ?", id).First(&thermograph).Error
	return
}

// GetThermographInfoList 分页获取Thermograph记录
// Author [piexlmax](https://github.com/piexlmax)
func (thermographService *ThermographService)GetThermographInfoList(info EpidemicInfoReq.ThermographSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Thermograph{})
    var thermographs []EpidemicInfo.Thermograph
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Area_id != nil {
        db = db.Where("area_id = ?",info.Area_id)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&thermographs).Error
	return err, thermographs, total
}
