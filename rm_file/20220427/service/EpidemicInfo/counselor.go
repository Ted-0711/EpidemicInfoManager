package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type CounselorService struct {
}

// CreateCounselor 创建Counselor记录
// Author [piexlmax](https://github.com/piexlmax)
func (counselorService *CounselorService) CreateCounselor(counselor EpidemicInfo.Counselor) (err error) {
	err = global.GVA_DB.Create(&counselor).Error
	return err
}

// DeleteCounselor 删除Counselor记录
// Author [piexlmax](https://github.com/piexlmax)
func (counselorService *CounselorService)DeleteCounselor(counselor EpidemicInfo.Counselor) (err error) {
	err = global.GVA_DB.Delete(&counselor).Error
	return err
}

// DeleteCounselorByIds 批量删除Counselor记录
// Author [piexlmax](https://github.com/piexlmax)
func (counselorService *CounselorService)DeleteCounselorByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Counselor{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCounselor 更新Counselor记录
// Author [piexlmax](https://github.com/piexlmax)
func (counselorService *CounselorService)UpdateCounselor(counselor EpidemicInfo.Counselor) (err error) {
	err = global.GVA_DB.Save(&counselor).Error
	return err
}

// GetCounselor 根据id获取Counselor记录
// Author [piexlmax](https://github.com/piexlmax)
func (counselorService *CounselorService)GetCounselor(id uint) (err error, counselor EpidemicInfo.Counselor) {
	err = global.GVA_DB.Where("id = ?", id).First(&counselor).Error
	return
}

// GetCounselorInfoList 分页获取Counselor记录
// Author [piexlmax](https://github.com/piexlmax)
func (counselorService *CounselorService)GetCounselorInfoList(info EpidemicInfoReq.CounselorSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Counselor{})
    var counselors []EpidemicInfo.Counselor
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Cnslr_id != "" {
        db = db.Where("cnslr_id LIKE ?","%"+ info.Cnslr_id+"%")
    }
    if info.Dept_name != nil {
        db = db.Where("dept_name = ?",info.Dept_name)
    }
    if info.Cnslr_name != "" {
        db = db.Where("cnslr_name LIKE ?","%"+ info.Cnslr_name+"%")
    }
    if info.Cnslr_gender != nil {
        db = db.Where("cnslr_gender = ?",info.Cnslr_gender)
    }
    if info.Cnslr_phone_number != "" {
        db = db.Where("cnslr_phone_number LIKE ?","%"+ info.Cnslr_phone_number+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&counselors).Error
	return err, counselors, total
}
