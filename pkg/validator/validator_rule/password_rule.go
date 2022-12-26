package validatorrule

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"sync"
)

type PasswordRule struct{ *CommonRule }

func NewPasswordRule() *PasswordRule {
	return &PasswordRule{&CommonRule{
		Field:     "password",
		FieldName: "Password",
		Message:   "{0} invalid format",
	}}
}

var (
	regexOnce                                sync.Once
	checkChar, needNum, needLower, needUpper *regexp.Regexp
)

func (*PasswordRule) GetRule() func(fl validator.FieldLevel) bool {
	regexOnce.Do(func() {
		checkChar, _ = regexp.Compile(`^[A-Za-z\d@$!%*#?&]{6,32}$`)
		needNum, _ = regexp.Compile(`\d`)
		needLower, _ = regexp.Compile(`[a-z]`)
		needUpper, _ = regexp.Compile(`[A-Z]`)
	})
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().String()
		if data != "" {
			// JS used ^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)[A-Za-z\d@$!%*#?&]{8,32}$
			if checkChar.MatchString(data) &&
				needNum.MatchString(data) &&
				needLower.MatchString(data) &&
				needUpper.MatchString(data) {
				return true
			}
		}
		return false
	}
}
