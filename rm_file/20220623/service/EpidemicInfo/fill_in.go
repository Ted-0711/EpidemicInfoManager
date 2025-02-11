package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type Fill_inService struct {
}

// CreateFill_in 创建Fill_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (fill_inService *Fill_inService) CreateFill_in(fill_in EpidemicInfo.Fill_in) (err error) {
	err = global.GVA_DB.Create(&fill_in).Error
	return err
}

// DeleteFill_in 删除Fill_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (fill_inService *Fill_inService)DeleteFill_in(fill_in EpidemicInfo.Fill_in) (err error) {
	err = global.GVA_DB.Delete(&fill_in).Error
	return err
}

// DeleteFill_inByIds 批量删除Fill_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (fill_inService *Fill_inService)DeleteFill_inByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Fill_in{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFill_in 更新Fill_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (fill_inService *Fill_inService)UpdateFill_in(fill_in EpidemicInfo.Fill_in) (err error) {
	err = global.GVA_DB.Save(&fill_in).Error
	return err
}

// GetFill_in 根据id获取Fill_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (fill_inService *Fill_inService)GetFill_in(id uint) (err error, fill_in EpidemicInfo.Fill_in) {
	err = global.GVA_DB.Where("id = ?", id).First(&fill_in).Error
	return
}

// GetFill_inInfoList 分页获取Fill_in记录
// Author [piexlmax](https://github.com/piexlmax)
func (fill_inService *Fill_inService)GetFill_inInfoList(info EpidemicInfoReq.Fill_inSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Fill_in{})
    var fill_ins []EpidemicInfo.Fill_in
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Student_id != "" {
        db = db.Where("student_id LIKE ?","%"+ info.Student_id+"%")
    }
    if info.Qtn_id != nil {
        db = db.Where("qtn_id = ?",info.Qtn_id)
    }
    if info.Fill_in_content != "" {
        db = db.Where("fill_in_content LIKE ?","%"+ info.Fill_in_content+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&fill_ins).Error
	return err, fill_ins, total
}
