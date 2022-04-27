package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type LabService struct {
}

// CreateLab 创建Lab记录
// Author [piexlmax](https://github.com/piexlmax)
func (labService *LabService) CreateLab(lab EpidemicInfo.Lab) (err error) {
	err = global.GVA_DB.Create(&lab).Error
	return err
}

// DeleteLab 删除Lab记录
// Author [piexlmax](https://github.com/piexlmax)
func (labService *LabService)DeleteLab(lab EpidemicInfo.Lab) (err error) {
	err = global.GVA_DB.Delete(&lab).Error
	return err
}

// DeleteLabByIds 批量删除Lab记录
// Author [piexlmax](https://github.com/piexlmax)
func (labService *LabService)DeleteLabByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Lab{},"id in ?",ids.Ids).Error
	return err
}

// UpdateLab 更新Lab记录
// Author [piexlmax](https://github.com/piexlmax)
func (labService *LabService)UpdateLab(lab EpidemicInfo.Lab) (err error) {
	err = global.GVA_DB.Save(&lab).Error
	return err
}

// GetLab 根据id获取Lab记录
// Author [piexlmax](https://github.com/piexlmax)
func (labService *LabService)GetLab(id uint) (err error, lab EpidemicInfo.Lab) {
	err = global.GVA_DB.Where("id = ?", id).First(&lab).Error
	return
}

// GetLabInfoList 分页获取Lab记录
// Author [piexlmax](https://github.com/piexlmax)
func (labService *LabService)GetLabInfoList(info EpidemicInfoReq.LabSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Lab{})
    var labs []EpidemicInfo.Lab
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Area_id != nil {
        db = db.Where("area_id = ?",info.Area_id)
    }
    if info.Lab_name != "" {
        db = db.Where("lab_name LIKE ?","%"+ info.Lab_name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&labs).Error
	return err, labs, total
}
