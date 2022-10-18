package entity

type Person struct {
	Id           string
	Login        string
	Password     string
	HashPassword []byte
}
