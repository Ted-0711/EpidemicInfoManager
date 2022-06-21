// 自动生成模板Quarantine
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Quarantine 结构体
// 如果含有time.Time 请自行import time包
type Quarantine struct {
	global.GVA_MODEL
	Student_id      string     `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:7;"`
	Quar_start_date *time.Time `json:"quar_start_date" form:"quar_start_date" gorm:"column:quar_start_date;comment:;"`
	Quar_site_id    *int       `json:"quar_site_id" form:"quar_site_id" gorm:"column:quar_site_id;comment:;"`
}

// TableName Quarantine 表名
func (Quarantine) TableName() string {
	return "epi_quarantine"
}
