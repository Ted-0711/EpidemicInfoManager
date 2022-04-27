// 自动生成模板Fill_in
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Fill_in 结构体
// 如果含有time.Time 请自行import time包
type Fill_in struct {
	global.GVA_MODEL
	Qtn_id          *int       `json:"qtn_id" form:"qtn_id" gorm:"column:qtn_id;comment:;"`
	Fill_in_time    *time.Time `json:"fill_in_time" form:"fill_in_time" gorm:"column:fill_in_time;comment:;"`
	Fill_in_content string     `json:"fill_in_content" form:"fill_in_content" gorm:"column:fill_in_content;comment:;"`
}

// TableName Fill_in 表名
func (Fill_in) TableName() string {
	return "epi_fill_in"
}
