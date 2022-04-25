// 自动生成模板Thermograph
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Thermograph 结构体
// 如果含有time.Time 请自行import time包
type Thermograph struct {
      global.GVA_MODEL
      Therm_id  string `json:"therm_id" form:"therm_id" gorm:"column:therm_id;comment:;"`
      Area_id  string `json:"area_id" form:"area_id" gorm:"column:area_id;comment:;"`
}


// TableName Thermograph 表名
func (Thermograph) TableName() string {
  return "thermograph"
}

