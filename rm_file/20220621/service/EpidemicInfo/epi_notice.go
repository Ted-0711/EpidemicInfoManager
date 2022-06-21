package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type NoticeService struct {
}

// CreateNotice 创建Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService) CreateNotice(notice EpidemicInfo.Notice) (err error) {
	err = global.GVA_DB.Create(&notice).Error
	return err
}

// DeleteNotice 删除Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService)DeleteNotice(notice EpidemicInfo.Notice) (err error) {
	err = global.GVA_DB.Delete(&notice).Error
	return err
}

// DeleteNoticeByIds 批量删除Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService)DeleteNoticeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Notice{},"id in ?",ids.Ids).Error
	return err
}

// UpdateNotice 更新Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService)UpdateNotice(notice EpidemicInfo.Notice) (err error) {
	err = global.GVA_DB.Save(&notice).Error
	return err
}

// GetNotice 根据id获取Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService)GetNotice(id uint) (err error, notice EpidemicInfo.Notice) {
	err = global.GVA_DB.Where("id = ?", id).First(&notice).Error
	return
}

// GetNoticeInfoList 分页获取Notice记录
// Author [piexlmax](https://github.com/piexlmax)
func (noticeService *NoticeService)GetNoticeInfoList(info EpidemicInfoReq.NoticeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Notice{})
    var notices []EpidemicInfo.Notice
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Title != "" {
        db = db.Where("title LIKE ?","%"+ info.Title+"%")
    }
    if info.Content != "" {
        db = db.Where("content LIKE ?","%"+ info.Content+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&notices).Error
	return err, notices, total
}
