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
	Manufacturer   *int       `json:"manufacturer" form:"manufacturer" gorm:"column:manufacturer;comment:;"`
	Vaccine_type   *int       `json:"vaccine_type" form:"vaccine_type" gorm:"column:vaccine_type;comment:;"`
	Inoculate_date *time.Time `json:"inoculate_date" form:"inoculate_date" gorm:"column:inoculate_date;comment:;"`
	Screenshot_url string     `json:"screenshot_url" form:"screenshot_url" gorm:"column:screenshot_url;comment:;size:70;"`
}

// TableName Vaccine 表名
func (Vaccine) TableName() string {
	return "epi_vaccine"
}
