// 自动生成模板Questionnaire
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Questionnaire 结构体
// 如果含有time.Time 请自行import time包
type Questionnaire struct {
	global.GVA_MODEL
	Qtn_title    string     `json:"qtn_title" form:"qtn_title" gorm:"column:qtn_title;comment:;"`
	Qtn_deadline *time.Time `json:"qtn_deadline" form:"qtn_deadline" gorm:"column:qtn_deadline;comment:;"`
	Qtn_q1       string     `json:"qtn_q1" form:"qtn_q1" gorm:"column:qtn_q1;comment:;"`
	Qtn_q2       string     `json:"qtn_q2" form:"qtn_q2" gorm:"column:qtn_q2;comment:;"`
	Qtn_q3       string     `json:"qtn_q3" form:"qtn_q3" gorm:"column:qtn_q3;comment:;"`
	Qtn_q4       string     `json:"qtn_q4" form:"qtn_q4" gorm:"column:qtn_q4;comment:;"`
	Qtn_q5       string     `json:"qtn_q5" form:"qtn_q5" gorm:"column:qtn_q5;comment:;"`
}

// TableName Questionnaire 表名
func (Questionnaire) TableName() string {
	return "epi_questionnaire"
}
