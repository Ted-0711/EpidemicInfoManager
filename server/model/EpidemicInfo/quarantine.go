// 自动生成模板Quarantine
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Quarantine 结构体
// 如果含有time.Time 请自行import time包
type Quarantine struct {
      global.GVA_MODEL
      Student_id  string `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:7;"`
      Quar_start_date  *time.Time `json:"quar_start_date" form:"quar_start_date" gorm:"column:quar_start_date;comment:;"`
      Quar_end_date  *time.Time `json:"quar_end_date" form:"quar_end_date" gorm:"column:quar_end_date;comment:;"`
      Quar_site  string `json:"quar_site" form:"quar_site" gorm:"column:quar_site;comment:;size:100;"`
}


// TableName Quarantine 表名
func (Quarantine) TableName() string {
  return "epi_quarantine"
}

