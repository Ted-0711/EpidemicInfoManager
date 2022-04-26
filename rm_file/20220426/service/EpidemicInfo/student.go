package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type StudentService struct {
}

// CreateStudent 创建Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) CreateStudent(student EpidemicInfo.Student) (err error) {
	err = global.GVA_DB.Create(&student).Error
	return err
}

// DeleteStudent 删除Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService)DeleteStudent(student EpidemicInfo.Student) (err error) {
	err = global.GVA_DB.Delete(&student).Error
	return err
}

// DeleteStudentByIds 批量删除Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService)DeleteStudentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Student{},"id in ?",ids.Ids).Error
	return err
}

// UpdateStudent 更新Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService)UpdateStudent(student EpidemicInfo.Student) (err error) {
	err = global.GVA_DB.Save(&student).Error
	return err
}

// GetStudent 根据id获取Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService)GetStudent(id uint) (err error, student EpidemicInfo.Student) {
	err = global.GVA_DB.Where("id = ?", id).First(&student).Error
	return
}

// GetStudentInfoList 分页获取Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService)GetStudentInfoList(info EpidemicInfoReq.StudentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Student{})
    var students []EpidemicInfo.Student
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Student_id != "" {
        db = db.Where("student_id = ?",info.Student_id)
    }
    if info.Dept_name != nil {
        db = db.Where("dept_name = ?",info.Dept_name)
    }
    if info.Student_name != "" {
        db = db.Where("student_name LIKE ?","%"+ info.Student_name+"%")
    }
    if info.Student_gender != nil {
        db = db.Where("student_gender = ?",info.Student_gender)
    }
    if info.Student_phone_number != "" {
        db = db.Where("student_phone_number LIKE ?","%"+ info.Student_phone_number+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&students).Error
	return err, students, total
}
