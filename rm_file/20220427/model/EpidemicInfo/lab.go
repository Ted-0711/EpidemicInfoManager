// 自动生成模板Lab
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Lab 结构体
// 如果含有time.Time 请自行import time包
type Lab struct {
      global.GVA_MODEL
      Area_id  *int `json:"area_id" form:"area_id" gorm:"column:area_id;comment:;"`
      Lab_name  string `json:"lab_name" form:"lab_name" gorm:"column:lab_name;comment:;"`
}


// TableName Lab 表名
func (Lab) TableName() string {
  return "epi_lab"
}

