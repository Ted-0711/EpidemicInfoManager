// 自动生成模板Clock_in
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Clock_in 结构体
// 如果含有time.Time 请自行import time包
type Clock_in struct {
	global.GVA_MODEL
	Student_id    string     `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:7;"`
	Clock_in_date *time.Time `json:"clock_in_date" form:"clock_in_date" gorm:"column:clock_in_date;comment:;"`
	Area_id       *int       `json:"area_id" form:"area_id" gorm:"column:area_id;comment:;"`
	Temperature   *float64   `json:"temperature" form:"temperature" gorm:"column:temperature;comment:;"`
	Symptom       string     `json:"symptom" form:"symptom" gorm:"column:symptom;comment:;"`
}

// TableName Clock_in 表名
func (Clock_in) TableName() string {
	return "epi_clock_in"
}
