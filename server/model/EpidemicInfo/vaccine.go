// 自动生成模板Vaccine
package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Vaccine 结构体
// 如果含有time.Time 请自行import time包
type Vaccine struct {
	global.GVA_MODEL
	Student_id     string     `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;size:7;"`
	Mfr_id         *int       `json:"mfr_id" form:"mfr_id" gorm:"column:mfr_id;comment:;"`
	Inoculate_date *time.Time `json:"inoculate_date" form:"inoculate_date" gorm:"column:inoculate_date;comment:;"`
	Prod_date      *time.Time `json:"prod_date" form:"prod_date" gorm:"column:prod_date;comment:;"`
}

// TableName Vaccine 表名
func (Vaccine) TableName() string {
	return "epi_vaccine"
}
