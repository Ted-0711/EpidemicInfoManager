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
	Qtn_content  string     `json:"qtn_content" form:"qtn_content" gorm:"column:qtn_content;comment:;"`
	Qtn_deadline *time.Time `json:"qtn_deadline" form:"qtn_deadline" gorm:"column:qtn_deadline;comment:;"`
}

// TableName Questionnaire 表名
func (Questionnaire) TableName() string {
	return "epi_questionnaire"
}
