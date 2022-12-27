package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"models/status_code"
	"reflect"
	"shareerrors"
	"strings"
	vr "validator/validator_rule"
)

var (
	uni      *ut.UniversalTranslator
	trans    ut.Translator
	Validate *validator.Validate
)

func init() {
	en := en.New()
	uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ = uni.GetTranslator("en")

	Validate = validator.New()
	en_translations.RegisterDefaultTranslations(Validate, trans)

	//vr.RegisterValidationRule(vr.NewPhoneRule(), Validate, trans)
	vr.RegisterValidationRule(vr.NewPasswordRule(), Validate, trans)
	//vr.RegisterValidationRule(vr.NewUsernameRule(), Validate, trans)
	//vr.RegisterValidationRule(vr.NewBankAccountRule(), Validate, trans)
}

func ValidateStruct[T any](structRule T) *shareerrors.Error {
	if err := Validate.Struct(structRule); err != nil {
		formError := shareerrors.NewError(status_code.BadRequest, "Wrong Input")
		if err != nil {
			fieldError := make([]shareerrors.FieldError, len(err.(validator.ValidationErrors)))
			for i, e := range err.(validator.ValidationErrors) {

				jsonFieldName := e.Field()
				if field, ok := reflect.TypeOf(structRule).Elem().FieldByName(e.Field()); ok {
					if jsonTag, ok := field.Tag.Lookup("json"); ok {
						println(jsonTag)
						jsonFieldName = strings.Split(jsonTag, ",")[0]
					}
				}

				fieldError[i] = shareerrors.FieldError{
					FieldName:   jsonFieldName,
					Description: e.Translate(trans),
				}
			}
			formError.Data = fieldError
		}

		return formError
	}

	return nil

}

func ValidateVar(fieldName string, err error) error {
	if err != nil {
		formError := shareerrors.NewError(status_code.BadRequest, "Wrong Input")
		fieldError := make([]shareerrors.FieldError, len(err.(validator.ValidationErrors)))

		for i, e := range err.(validator.ValidationErrors) {
			fieldError[i] = shareerrors.FieldError{
				FieldName:   fieldName,
				Description: e.Translate(trans),
			}
		}

		formError.Data = fieldError

		return formError
	}
	return nil
}
