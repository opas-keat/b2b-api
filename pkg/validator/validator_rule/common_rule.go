package validatorrule

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Rule interface {
	RegisterMessageTranslate(ut.Translator) error
	RegisterFieldNameTranslate(ut.Translator, validator.FieldError) string
	GetField() string
	GetFieldName() string
	GetMessage() string
	GetRule() func(fl validator.FieldLevel) bool
}

type CommonRule struct {
	Field     string
	FieldName string
	Message   string
}

func (rule *CommonRule) GetField() string {
	return rule.Field
}

func (rule *CommonRule) GetFieldName() string {
	return rule.FieldName
}

func (rule *CommonRule) GetMessage() string {
	return rule.Message
}

func (rule *CommonRule) RegisterMessageTranslate(ut ut.Translator) error {
	return ut.Add(rule.Field, rule.Message, true)
}

func (rule *CommonRule) RegisterFieldNameTranslate(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T(rule.Field, rule.FieldName)
	return t
}

func RegisterValidationRule(rule Rule, validate *validator.Validate, trans ut.Translator) {
	_ = validate.RegisterValidation(rule.GetField(), rule.GetRule())
	_ = validate.RegisterTranslation(rule.GetField(), trans, rule.RegisterMessageTranslate, rule.RegisterFieldNameTranslate)
}
