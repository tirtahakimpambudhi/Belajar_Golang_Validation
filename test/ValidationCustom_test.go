package test

import (
	"github.com/go-playground/validator/v10"
	"go_validation/internal/domain/model"
	"go_validation/internal/exception"
	"go_validation/internal/validation"
	"testing"
)

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validation.NewCustomValidator(validate)
	validate.RegisterAlias("varchar", "required,max=255")
	validate.RegisterAlias("nis", "required")

	student := model.NewStudent(34672, "Tirta Hakim Pambudhi", "XI SIJA 2", "2006", "+63 1232 4562 8890")
	if err := validate.Struct(student); err != nil {
		validationnErr := err.(validator.ValidationErrors)
		for _, fieldError := range validationnErr {
			t.Log(exception.NewErrorMessage(fieldError.Field(), fieldError.Value(), fieldError.Tag()).Error())
		}
	}
}

func TestCustomValdationCrossField(t *testing.T) {
	validate := validator.New()
	validation.NewCustomValidator(validate)
	validate.RegisterAlias("varchar", "required,max=255")
	validate.RegisterAlias("password", "required,min=8")

	user := model.NewLoginUsers("test@gmail.com", "test@gmail.com", "+62 8421 2133 1234", "12345678", "12345678")
	if err := validate.Struct(user); err != nil {
		validationnErr := err.(validator.ValidationErrors)
		for _, fieldError := range validationnErr {
			t.Log(exception.NewErrorMessage(fieldError.Field(), fieldError.Value(), fieldError.Tag()).Error())
		}
	}
}
