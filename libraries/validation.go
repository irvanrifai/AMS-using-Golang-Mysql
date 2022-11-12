package libraries

import (
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

type Validation struct {
	validate *validator.Validate
    translator ut.Translator
}

func NewValidation() *Validation {
    translator := en.New()
    uni := ut.New(translator, translator)

    trans, _ := uni.GetTranslator("ën")

    validate := validator.New()

    en_translation.RegisterDefaultTranslations(validate, trans)

    // register tag label (name field custom)
    validate.RegisterTagNameFunc(func(field reflect.StructField) string {
        name := field.Tag.Get("label")
        return name
    })

    // make errorr custom message
    validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
        return ut.Add("required", "{0} is cannot null", true)
    }, func(ut ut.Translator, fe validator.FieldError) string {
        t,_ := ut.T("required", fe.Field())
        return t
    })

    return &Validation{
        validate: validate,
        translator: trans,
    }
}

func (v *Validation) Struct(s interface{}) interface{}  {
    errors := make(map[string]string)

    err := v.validate.Struct(s)
    if err != nil {
        for _, e := range err.(validator.ValidationErrors){
            errors[e.StructField()] = e.Translate(v.translator)
        }
    }

    if len(errors) > 0 {
        return errors
    }

    return nil
}