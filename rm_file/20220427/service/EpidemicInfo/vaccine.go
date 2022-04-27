package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type VaccineService struct {
}

// CreateVaccine 创建Vaccine记录
// Author [piexlmax](https://github.com/piexlmax)
func (vaccineService *VaccineService) CreateVaccine(vaccine EpidemicInfo.Vaccine) (err error) {
	err = global.GVA_DB.Create(&vaccine).Error
	return err
}

// DeleteVaccine 删除Vaccine记录
// Author [piexlmax](https://github.com/piexlmax)
func (vaccineService *VaccineService)DeleteVaccine(vaccine EpidemicInfo.Vaccine) (err error) {
	err = global.GVA_DB.Delete(&vaccine).Error
	return err
}

// DeleteVaccineByIds 批量删除Vaccine记录
// Author [piexlmax](https://github.com/piexlmax)
func (vaccineService *VaccineService)DeleteVaccineByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Vaccine{},"id in ?",ids.Ids).Error
	return err
}

// UpdateVaccine 更新Vaccine记录
// Author [piexlmax](https://github.com/piexlmax)
func (vaccineService *VaccineService)UpdateVaccine(vaccine EpidemicInfo.Vaccine) (err error) {
	err = global.GVA_DB.Save(&vaccine).Error
	return err
}

// GetVaccine 根据id获取Vaccine记录
// Author [piexlmax](https://github.com/piexlmax)
func (vaccineService *VaccineService)GetVaccine(id uint) (err error, vaccine EpidemicInfo.Vaccine) {
	err = global.GVA_DB.Where("id = ?", id).First(&vaccine).Error
	return
}

// GetVaccineInfoList 分页获取Vaccine记录
// Author [piexlmax](https://github.com/piexlmax)
func (vaccineService *VaccineService)GetVaccineInfoList(info EpidemicInfoReq.VaccineSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Vaccine{})
    var vaccines []EpidemicInfo.Vaccine
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Student_id != "" {
        db = db.Where("student_id LIKE ?","%"+ info.Student_id+"%")
    }
    if info.Mfr_id != nil {
        db = db.Where("mfr_id = ?",info.Mfr_id)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&vaccines).Error
	return err, vaccines, total
}
