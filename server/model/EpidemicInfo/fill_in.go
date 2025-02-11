// 自动生成模板Fill_in
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Fill_in 结构体
// 如果含有time.Time 请自行import time包
type Fill_in struct {
      global.GVA_MODEL
      Student_id  string `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:7;"`
      Qtn_id  *int `json:"qtn_id" form:"qtn_id" gorm:"column:qtn_id;comment:;"`
      Fill_in_a1  string `json:"fill_in_a1" form:"fill_in_a1" gorm:"column:fill_in_a1;comment:;"`
      Fill_in_a2  string `json:"fill_in_a2" form:"fill_in_a2" gorm:"column:fill_in_a2;comment:;"`
      Fill_in_a3  string `json:"fill_in_a3" form:"fill_in_a3" gorm:"column:fill_in_a3;comment:;"`
      Fill_in_a4  string `json:"fill_in_a4" form:"fill_in_a4" gorm:"column:fill_in_a4;comment:;"`
      Fill_in_a5  string `json:"fill_in_a5" form:"fill_in_a5" gorm:"column:fill_in_a5;comment:;"`
}


// TableName Fill_in 表名
func (Fill_in) TableName() string {
  return "epi_fill_in"
}

