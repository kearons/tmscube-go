package dao

import "tmscube-go/common/model"

type TemplateDao struct {}

func (d *TemplateDao) GetTemplateByKey(templateKey string) (model.TemplateDicModel, error) {
	var (
		t model.TemplateDicModel
	)

	r := t.DB().Where("template_key = ?", templateKey).First(&t)

	return t, r.Error
}