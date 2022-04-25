// 自动生成模板Measure
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Measure 结构体
// 如果含有time.Time 请自行import time包
type Measure struct {
      global.GVA_MODEL
      Student_id  string `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:7;"`
}


// TableName Measure 表名
func (Measure) TableName() string {
  return "measure"
}

