package model

type User struct {
	ID     string      `json:"id"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Person *UserPerson `json:"person"`
}

type UserPerson struct {
	Email string `json:"email"`
}
