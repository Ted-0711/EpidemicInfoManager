package EpidemicInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    EpidemicInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo/request"
)

type QuestionnaireService struct {
}

// CreateQuestionnaire 创建Questionnaire记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService) CreateQuestionnaire(questionnaire EpidemicInfo.Questionnaire) (err error) {
	err = global.GVA_DB.Create(&questionnaire).Error
	return err
}

// DeleteQuestionnaire 删除Questionnaire记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService)DeleteQuestionnaire(questionnaire EpidemicInfo.Questionnaire) (err error) {
	err = global.GVA_DB.Delete(&questionnaire).Error
	return err
}

// DeleteQuestionnaireByIds 批量删除Questionnaire记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService)DeleteQuestionnaireByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]EpidemicInfo.Questionnaire{},"id in ?",ids.Ids).Error
	return err
}

// UpdateQuestionnaire 更新Questionnaire记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService)UpdateQuestionnaire(questionnaire EpidemicInfo.Questionnaire) (err error) {
	err = global.GVA_DB.Save(&questionnaire).Error
	return err
}

// GetQuestionnaire 根据id获取Questionnaire记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService)GetQuestionnaire(id uint) (err error, questionnaire EpidemicInfo.Questionnaire) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionnaire).Error
	return
}

// GetQuestionnaireInfoList 分页获取Questionnaire记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService)GetQuestionnaireInfoList(info EpidemicInfoReq.QuestionnaireSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&EpidemicInfo.Questionnaire{})
    var questionnaires []EpidemicInfo.Questionnaire
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Qtn_title != "" {
        db = db.Where("qtn_title LIKE ?","%"+ info.Qtn_title+"%")
    }
    if info.Qtn_q1 != "" {
        db = db.Where("qtn_q1 LIKE ?","%"+ info.Qtn_q1+"%")
    }
    if info.Qtn_q2 != "" {
        db = db.Where("qtn_q2 LIKE ?","%"+ info.Qtn_q2+"%")
    }
    if info.Qtn_q3 != "" {
        db = db.Where("qtn_q3 LIKE ?","%"+ info.Qtn_q3+"%")
    }
    if info.Qtn_q4 != "" {
        db = db.Where("qtn_q4 LIKE ?","%"+ info.Qtn_q4+"%")
    }
    if info.Qtn_q5 != "" {
        db = db.Where("qtn_q5 LIKE ?","%"+ info.Qtn_q5+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&questionnaires).Error
	return err, questionnaires, total
}
