package service

import (
	"errors"
	"tmscube-go/common/model"
	"tmscube-go/constant"
)

type TemplateService struct{}

func (s *TemplateService) CompileTemplate(templateKey string, keywords map[string]string) (*model.TemplateDicModel, error) {

	template, err := s.getTemplateFromCache(templateKey)
	if err != nil {
		return nil, err
	}

	if _, ok := constant.DefaultPushType[template.MessageType]; !ok {
		return nil, errors.New("模板"+templateKey+":message_type配置错误")
	}

	if template.Status == 10 {
		return nil, errors.New("template {" + templateKey + "} banned")
	}

	if template.IsPhonePush == 0 {
		return nil, errors.New("template {" + templateKey + "} set forbiden")
	}

	if len(template.Keywords) == 0 {
		template.CompiledContent = template.Content
		return &template, nil
	}

	var (
		c   = 0
		tmp = ""
		r   = ""
	)
CompileLoop:
	for _, s := range template.Content {
		t := string(s)
		if t == "#" {
			c += 1
			if c < 4 {
				continue
			}
		} else if c == 1 || c == 3 {
			c = 0
		}

		switch c {
		case 0:
			r += t
		case 2:
			tmp += t
		case 4:
			if v, ok := keywords[tmp]; ok {
				r += v
				c = 0
				tmp = ""
			} else {
				err = errors.New("missing params:" + tmp)
				break CompileLoop
			}
		}
	}

	if err != nil {
		return nil, err
	}

	if c > 0 {
		return nil, errors.New("params is not replaced")
	}

	template.CompiledContent = r

	return &template, nil
}

func (s *TemplateService) getTemplateFromCache(templateKey string) (model.TemplateDicModel, error) {
	return templateDao.GetTemplateByKey(templateKey)
}
