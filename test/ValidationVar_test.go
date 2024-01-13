package test

import (
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en3 "github.com/go-playground/validator/v10/translations/en"
	"github.com/stretchr/testify/assert"
	"go_validation/helper"
	"go_validation/internal/exception"
	"testing"
)

func TestVarValidation(t *testing.T) {
	var customErr *exception.ErrorMessage
	validation := validator.New()
	input := ""
	err := validation.Var(input, "required")

	if err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			t.Fatal(err.Error())
		}
		for _, e := range validationErr {
			customErr = exception.NewErrorMessage(e.Field(), e.Value(), e.Tag())
		}
	}

	t.Log(customErr.Error())
}

func TestVarsValidation(t *testing.T) {
	validation := validator.New()
	input := ""
	input2 := ""
	err := validation.VarWithValue(input, input2, "required")

	password := "secret"
	confPass := "rahasia"

	errPass := validation.VarWithValue(password, confPass, "eqfield")

	assert.NotEqual(t, nil, err)
	assert.NotEqual(t, nil, errPass)
}

func TestVarWithTranslation(t *testing.T) {
	validation := validator.New()
	en := en2.New()
	uni := ut.New(en, en)
	trans, fn := uni.GetTranslator("en")
	if !fn {
		t.Fatal("Not Found Language")
	}
	en3.RegisterDefaultTranslations(validation, trans)

	input := ""
	tag := "required"

	validationErr := validation.Var(input, tag)

	errs := helper.TranslateError(validationErr, trans)
	t.Log(errs)
}
