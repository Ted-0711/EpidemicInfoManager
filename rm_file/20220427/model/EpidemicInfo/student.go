// 自动生成模板Student
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Student 结构体
// 如果含有time.Time 请自行import time包
type Student struct {
      global.GVA_MODEL
      Student_id  string `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:7;"`
      Dept_name  *int `json:"dept_name" form:"dept_name" gorm:"column:dept_name;comment:;"`
      Student_name  string `json:"student_name" form:"student_name" gorm:"column:student_name;comment:;"`
      Student_gender  *int `json:"student_gender" form:"student_gender" gorm:"column:student_gender;comment:;"`
      Student_phone_number  string `json:"student_phone_number" form:"student_phone_number" gorm:"column:student_phone_number;comment:;size:11;"`
}


// TableName Student 表名
func (Student) TableName() string {
  return "epi_student"
}

