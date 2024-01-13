package test

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go_validation/internal/domain/model"
	"testing"
	"time"
)

var GlobalFriends = model.NewRegisterUsers(uuid.NewString(), "testing_users", "123456", "123456")

func TestValidationStructBooks(t *testing.T) {
	validation := validator.New()
	now := time.Now()
	books := model.NewBook("978-602-8519-93-9", "Example Books", "Testify Books", "2024", false, &now, &now)

	err := validation.Struct(books)

	if err != nil {
		t.Log(err.Error())
	}

	t.Log(books)
}

func TestValidationStructUsers(t *testing.T) {
	validation := validator.New()
	id := uuid.NewString()
	user := model.NewRegisterUsers(id, "testing_users", "123456", "1234")

	err := validation.Struct(user)

	if err != nil {
		t.Log(err.Error())
	}

	t.Log(user)
}

func TestValidationStructTx(t *testing.T) {
	validation := validator.New()
	now := time.Now()
	id := uuid.NewString()
	tx := model.NewTransaction(id, float64(100), float64(150), float64(50), "No Description", &now, &now)

	err := validation.Struct(tx)

	if err != nil {
		t.Log(err.Error())
	}

	t.Log(tx)
}

func TestValidationDetailUser(t *testing.T) {
	friends := []model.RegisterUsers{*GlobalFriends}

	validation := validator.New()
	addr := model.Address{
		City:       "DKI JAKARTA",
		Country:    "ID",
		Street:     "JL. Jendral Soedirman",
		State:      "Jakarta",
		PostalCode: "52840",
	}

	detailUsers := model.NewDetailUsers(&addr, []string{"Mancing", "Bola", "Ngoding"}, friends)

	err := validation.Struct(detailUsers)

	if err != nil {
		t.Log(err.Error())
	}

	t.Log(detailUsers)
}

func TestValidationStructUsersAliasTag(t *testing.T) {
	validation := validator.New()
	id := uuid.NewString()
	validation.RegisterAlias("password", "required,min=8")
	wallet := map[string]int{
		"BCA":          100000000,
		"BRI":          400000000,
		"Bank Syariah": 500000000000000,
	}
	user := model.NewUsers(id, "testing_users", "12345678", "12345678", wallet)

	err := validation.Struct(user)

	if err != nil {
		t.Log(err.Error())
	}

	t.Log(user)
}
