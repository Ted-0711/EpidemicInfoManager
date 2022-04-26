package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type Quarantine_siteService struct {
}

// CreateQuarantine_site 创建Quarantine_site记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantine_siteService *Quarantine_siteService) CreateQuarantine_site(quarantine_site EpidemicInfo.Quarantine_site) (err error) {
	err = global.GVA_DB.Create(&quarantine_site).Error
	return err
}

// DeleteQuarantine_site 删除Quarantine_site记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantine_siteService *Quarantine_siteService)DeleteQuarantine_site(quarantine_site EpidemicInfo.Quarantine_site) (err error) {
	err = global.GVA_DB.Delete(&quarantine_site).Error
	return err
}

// DeleteQuarantine_siteByIds 批量删除Quarantine_site记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantine_siteService *Quarantine_siteService)DeleteQuarantine_siteByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Quarantine_site{},"id in ?",ids.Ids).Error
	return err
}

// UpdateQuarantine_site 更新Quarantine_site记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantine_siteService *Quarantine_siteService)UpdateQuarantine_site(quarantine_site EpidemicInfo.Quarantine_site) (err error) {
	err = global.GVA_DB.Save(&quarantine_site).Error
	return err
}

// GetQuarantine_site 根据id获取Quarantine_site记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantine_siteService *Quarantine_siteService)GetQuarantine_site(id uint) (err error, quarantine_site EpidemicInfo.Quarantine_site) {
	err = global.GVA_DB.Where("id = ?", id).First(&quarantine_site).Error
	return
}

// GetQuarantine_siteInfoList 分页获取Quarantine_site记录
// Author [piexlmax](https://github.com/piexlmax)
func (quarantine_siteService *Quarantine_siteService)GetQuarantine_siteInfoList(info EpidemicInfoReq.Quarantine_siteSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Quarantine_site{})
    var quarantine_sites []EpidemicInfo.Quarantine_site
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Quar_site_id != "" {
        db = db.Where("quar_site_id LIKE ?","%"+ info.Quar_site_id+"%")
    }
    if info.Area_id != "" {
        db = db.Where("area_id LIKE ?","%"+ info.Area_id+"%")
    }
    if info.Quar_site_name != "" {
        db = db.Where("quar_site_name LIKE ?","%"+ info.Quar_site_name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&quarantine_sites).Error
	return err, quarantine_sites, total
}
