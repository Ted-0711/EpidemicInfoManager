// 自动生成模板Quarantine_site
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Quarantine_site 结构体
// 如果含有time.Time 请自行import time包
type Quarantine_site struct {
      global.GVA_MODEL
      Area_id  *int `json:"area_id" form:"area_id" gorm:"column:area_id;comment:;"`
      Quar_site_name  string `json:"quar_site_name" form:"quar_site_name" gorm:"column:quar_site_name;comment:;"`
}


// TableName Quarantine_site 表名
func (Quarantine_site) TableName() string {
  return "epi_quarantine_site"
}

