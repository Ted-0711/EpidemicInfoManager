// 自动生成模板Temp_measure
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Temp_measure 结构体
// 如果含有time.Time 请自行import time包
type Temp_measure struct {
	global.GVA_MODEL
	Student_id        string     `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:7;"`
	Temp_measure_time *time.Time `json:"temp_measure_time" form:"temp_measure_time" gorm:"column:temp_measure_time;comment:;"`
	Therm_id          string     `json:"therm_id" form:"therm_id" gorm:"column:therm_id;comment:;size:5;"`
	Temperature       *float64   `json:"temperature" form:"temperature" gorm:"column:temperature;comment:;"`
}

// TableName Temp_measure 表名
func (Temp_measure) TableName() string {
	return "temp_measure"
}
