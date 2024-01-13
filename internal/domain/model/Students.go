package model

type Student struct {
	NIS       int    `json:"nis" validate:"nis,digit=5"`
	Name      string `json:"name" validate:"varchar"`
	Jurusan   string `json:"jurusan" validate:"varchar"`
	BirthDate string `json:"birthDate" validate:"required,validYear"`
	Phone     string `json:"phone" validate:"required,noWA"`
}

func NewStudent(NIS int, name string, jurusan string, birthDate string, phone string) *Student {
	return &Student{NIS: NIS, Name: name, Jurusan: jurusan, BirthDate: birthDate, Phone: phone}
}
