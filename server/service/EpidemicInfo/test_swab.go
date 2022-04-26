package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type Test_swabService struct {
}

// CreateTest_swab 创建Test_swab记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_swabService *Test_swabService) CreateTest_swab(test_swab EpidemicInfo.Test_swab) (err error) {
	err = global.GVA_DB.Create(&test_swab).Error
	return err
}

// DeleteTest_swab 删除Test_swab记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_swabService *Test_swabService)DeleteTest_swab(test_swab EpidemicInfo.Test_swab) (err error) {
	err = global.GVA_DB.Delete(&test_swab).Error
	return err
}

// DeleteTest_swabByIds 批量删除Test_swab记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_swabService *Test_swabService)DeleteTest_swabByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Test_swab{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTest_swab 更新Test_swab记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_swabService *Test_swabService)UpdateTest_swab(test_swab EpidemicInfo.Test_swab) (err error) {
	err = global.GVA_DB.Save(&test_swab).Error
	return err
}

// GetTest_swab 根据id获取Test_swab记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_swabService *Test_swabService)GetTest_swab(id uint) (err error, test_swab EpidemicInfo.Test_swab) {
	err = global.GVA_DB.Where("id = ?", id).First(&test_swab).Error
	return
}

// GetTest_swabInfoList 分页获取Test_swab记录
// Author [piexlmax](https://github.com/piexlmax)
func (test_swabService *Test_swabService)GetTest_swabInfoList(info EpidemicInfoReq.Test_swabSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Test_swab{})
    var test_swabs []EpidemicInfo.Test_swab
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Student_id != "" {
        db = db.Where("student_id LIKE ?","%"+ info.Student_id+"%")
    }
    if info.Lab_id != nil {
        db = db.Where("lab_id = ?",info.Lab_id)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&test_swabs).Error
	return err, test_swabs, total
}
