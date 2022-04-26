// 自动生成模板Migration
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Migration 结构体
// 如果含有time.Time 请自行import time包
type Migration struct {
	global.GVA_MODEL
	Student_id    string     `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:7;"`
	Start_area_id *int       `json:"start_area_id" form:"start_area_id" gorm:"column:start_area_id;comment:;"`
	Des_area_id   *int       `json:"des_area_id" form:"des_area_id" gorm:"column:des_area_id;comment:;"`
	Vehicle_type  *int       `json:"vehicle_type" form:"vehicle_type" gorm:"column:vehicle_type;comment:;"`
	Vehicle_info  string     `json:"vehicle_info" form:"vehicle_info" gorm:"column:vehicle_info;comment:;"`
	Mig_time      *time.Time `json:"mig_time" form:"mig_time" gorm:"column:mig_time;comment:;"`
}

// TableName Migration 表名
func (Migration) TableName() string {
	return "epi_migration"
}
