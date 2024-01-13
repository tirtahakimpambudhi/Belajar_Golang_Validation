package model

type RegisterUsers struct {
	ID              string `json:"id" validate:"required,uuid"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func NewRegisterUsers(ID string, username string, password string, confirmPassword string) *RegisterUsers {
	return &RegisterUsers{ID: ID, Username: username, Password: password, ConfirmPassword: confirmPassword}
}

type Address struct {
	Street     string `json:"street" validate:"required"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required"`
	Country    string `json:"country" validate:"required"`
}

type Hobbies []string

type DetailUsers struct {
	*Address `json:"address" validate:"required"`
	Hobbies  `json:"hobbies" validate:"required,len=3,dive,required,min=1"`
	Friends  []RegisterUsers `json:"friends" validate:"dive"`
}

func NewDetailUsers(address *Address, hobbies Hobbies, friends []RegisterUsers) *DetailUsers {
	return &DetailUsers{Address: address, Hobbies: hobbies, Friends: friends}
}

type Users struct {
	ID              string         `json:"id" validate:"required,uuid"`
	Username        string         `json:"username" validate:"required"`
	Password        string         `json:"password" validate:"password"`
	ConfirmPassword string         `json:"confirm_password" validate:"eqfield=Password,password"`
	Wallet          map[string]int `json:"wallet" validate:"dive,keys,required,min=2,endkeys,gt=0"`
}

func NewUsers(ID string, username string, password string, confirmPassword string, wallet map[string]int) *Users {
	return &Users{ID: ID, Username: username, Password: password, ConfirmPassword: confirmPassword, Wallet: wallet}
}

type LoginUsers struct {
	Username        string `json:"username" validate:"varchar,equals_fields=Email|equals_fields=Phone"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required,noWA"`
	Password        string `json:"password" validate:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"password,equals_fields=Password"`
}

func NewLoginUsers(username string, email string, phone string, password string, confirmPassword string) *LoginUsers {
	return &LoginUsers{Username: username, Email: email, Phone: phone, Password: password, ConfirmPassword: confirmPassword}
}
