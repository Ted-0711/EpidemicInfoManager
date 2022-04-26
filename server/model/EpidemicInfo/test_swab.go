// 自动生成模板Test_swab
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Test_swab 结构体
// 如果含有time.Time 请自行import time包
type Test_swab struct {
	global.GVA_MODEL
	Student_id  string     `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:7;"`
	Lab_id      *int       `json:"lab_id" form:"lab_id" gorm:"column:lab_id;comment:;"`
	Sample_date *time.Time `json:"sample_date" form:"sample_date" gorm:"column:sample_date;comment:;"`
	Test_date   *time.Time `json:"test_date" form:"test_date" gorm:"column:test_date;comment:;"`
	Test_result *int       `json:"test_result" form:"test_result" gorm:"column:test_result;comment:;"`
}

// TableName Test_swab 表名
func (Test_swab) TableName() string {
	return "epi_test_swab"
}
