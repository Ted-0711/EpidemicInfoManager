// 自动生成模板Area
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Area 结构体
// 如果含有time.Time 请自行import time包
type Area struct {
      global.GVA_MODEL
      Area_name  string `json:"area_name" form:"area_name" gorm:"column:area_name;comment:;"`
      Area_risk_level  *int `json:"area_risk_level" form:"area_risk_level" gorm:"column:area_risk_level;comment:;"`
}


// TableName Area 表名
func (Area) TableName() string {
  return "epi_area"
}

