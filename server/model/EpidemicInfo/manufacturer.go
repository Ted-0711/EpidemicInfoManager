// 自动生成模板Manufacturer
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Manufacturer 结构体
// 如果含有time.Time 请自行import time包
type Manufacturer struct {
      global.GVA_MODEL
      Area_id  *int `json:"area_id" form:"area_id" gorm:"column:area_id;comment:;"`
      Mfr_name  string `json:"mfr_name" form:"mfr_name" gorm:"column:mfr_name;comment:;"`
}


// TableName Manufacturer 表名
func (Manufacturer) TableName() string {
  return "epi_manufacturer"
}

