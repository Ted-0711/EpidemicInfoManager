package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type AreaService struct {
}

// CreateArea 创建Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) CreateArea(area EpidemicInfo.Area) (err error) {
	err = global.GVA_DB.Create(&area).Error
	return err
}

// DeleteArea 删除Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService)DeleteArea(area EpidemicInfo.Area) (err error) {
	err = global.GVA_DB.Delete(&area).Error
	return err
}

// DeleteAreaByIds 批量删除Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService)DeleteAreaByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Area{},"id in ?",ids.Ids).Error
	return err
}

// UpdateArea 更新Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService)UpdateArea(area EpidemicInfo.Area) (err error) {
	err = global.GVA_DB.Save(&area).Error
	return err
}

// GetArea 根据id获取Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService)GetArea(id uint) (err error, area EpidemicInfo.Area) {
	err = global.GVA_DB.Where("id = ?", id).First(&area).Error
	return
}

// GetAreaInfoList 分页获取Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService)GetAreaInfoList(info EpidemicInfoReq.AreaSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Area{})
    var areas []EpidemicInfo.Area
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Area_name != "" {
        db = db.Where("area_name LIKE ?","%"+ info.Area_name+"%")
    }
    if info.Area_location != "" {
        db = db.Where("area_location LIKE ?","%"+ info.Area_location+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&areas).Error
	return err, areas, total
}
