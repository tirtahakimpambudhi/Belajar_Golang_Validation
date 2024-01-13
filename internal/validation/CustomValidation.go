package validation

import (
	"github.com/go-playground/validator/v10"
	"go_validation/internal/domain/model"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func NewCustomValidator(validator *validator.Validate) {
	validator.RegisterValidation("noWA", IsValidWhatsAppNumber)
	validator.RegisterValidation("validYear", IsValidYear)
	validator.RegisterValidation("digit", CountDigit)
	validator.RegisterValidation("equals_fields", MustEqualsIgnoreCase)
	validator.RegisterStructValidation(UserValidation, model.LoginUsers{})
}

func IsValidYear(fl validator.FieldLevel) bool {
	year, err := time.Parse("2006", fl.Field().String())
	if err != nil {
		return false
	}

	return year.Year() >= 1900
}

func IsValidWhatsAppNumber(fl validator.FieldLevel) bool {
	number := fl.Field().String()
	regexPattern := `\+\d{1,3} \d{4} \d{4} \d{4}$`

	// Creating a regex object
	regex := regexp.MustCompile(regexPattern)

	// Checking if the number matches the regex pattern
	return regex.MatchString(number)
}

func CountDigit(fl validator.FieldLevel) bool {
	value := fl.Field().Int()
	valueStr := strconv.Itoa(int(value))

	length, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}
	return length == len(valueStr)
}

func MustEqualsIgnoreCase(fl validator.FieldLevel) bool {
	value, _, _, ok := fl.GetStructFieldOK2()
	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(value.String())
	secondValue := strings.ToUpper(fl.Field().String())
	return firstValue == secondValue
}

func UserValidation(level validator.StructLevel) {
	request := level.Current().Interface().(model.LoginUsers)

	if request.Username == request.Email || request.Username == request.Phone {

	} else {
		level.ReportError(request.Username, "Username", "Username", "username", "")
	}
}
