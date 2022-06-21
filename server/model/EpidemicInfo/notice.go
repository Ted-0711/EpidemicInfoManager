// 自动生成模板Notice
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Notice 结构体
// 如果含有time.Time 请自行import time包
type Notice struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;size:50;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:;size:200;"`
}


// TableName Notice 表名
func (Notice) TableName() string {
  return "epi_notice"
}

