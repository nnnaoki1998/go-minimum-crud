package model

type (
	UserId    int16
	UserName  string
	UserEmail string
)

type User struct {
	Id        UserId    `json:"id"`
	UserName  UserName  `json:"name"`
	UserEmail UserEmail `json:"email"`
}

type NewUser struct {
	UserName  UserName
	UserEmail UserEmail
}
