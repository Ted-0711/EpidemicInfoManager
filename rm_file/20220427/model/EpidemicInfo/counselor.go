// 自动生成模板Counselor
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Counselor 结构体
// 如果含有time.Time 请自行import time包
type Counselor struct {
      global.GVA_MODEL
      Cnslr_id  string `json:"cnslr_id" form:"cnslr_id" gorm:"column:cnslr_id;comment:;size:5;"`
      Dept_name  *int `json:"dept_name" form:"dept_name" gorm:"column:dept_name;comment:;"`
      Cnslr_name  string `json:"cnslr_name" form:"cnslr_name" gorm:"column:cnslr_name;comment:;"`
      Cnslr_gender  *int `json:"cnslr_gender" form:"cnslr_gender" gorm:"column:cnslr_gender;comment:;"`
      Cnslr_phone_number  string `json:"cnslr_phone_number" form:"cnslr_phone_number" gorm:"column:cnslr_phone_number;comment:;size:11;"`
}


// TableName Counselor 表名
func (Counselor) TableName() string {
  return "epi_counselor"
}

