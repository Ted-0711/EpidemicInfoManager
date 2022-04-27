package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type Clock_inService struct {
}

// CreateClock_in 创建Clock_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (clock_inService *Clock_inService) CreateClock_in(clock_in EpidemicInfo.Clock_in) (err error) {
	err = global.GVA_DB.Create(&clock_in).Error
	return err
}

// DeleteClock_in 删除Clock_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (clock_inService *Clock_inService)DeleteClock_in(clock_in EpidemicInfo.Clock_in) (err error) {
	err = global.GVA_DB.Delete(&clock_in).Error
	return err
}

// DeleteClock_inByIds 批量删除Clock_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (clock_inService *Clock_inService)DeleteClock_inByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Clock_in{},"id in ?",ids.Ids).Error
	return err
}

// UpdateClock_in 更新Clock_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (clock_inService *Clock_inService)UpdateClock_in(clock_in EpidemicInfo.Clock_in) (err error) {
	err = global.GVA_DB.Save(&clock_in).Error
	return err
}

// GetClock_in 根据id获取Clock_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (clock_inService *Clock_inService)GetClock_in(id uint) (err error, clock_in EpidemicInfo.Clock_in) {
	err = global.GVA_DB.Where("id = ?", id).First(&clock_in).Error
	return
}

// GetClock_inInfoList 分页获取Clock_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (clock_inService *Clock_inService)GetClock_inInfoList(info EpidemicInfoReq.Clock_inSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Clock_in{})
    var clock_ins []EpidemicInfo.Clock_in
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Student_id != "" {
        db = db.Where("student_id LIKE ?","%"+ info.Student_id+"%")
    }
    if info.Area_id != nil {
        db = db.Where("area_id = ?",info.Area_id)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&clock_ins).Error
	return err, clock_ins, total
}
