package models

type Person struct {
	Name     string `json:"name" xml:"name" form:"name"`
	Surname  string `json:"surname" xml:"surname" form:"surname"`
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
}
