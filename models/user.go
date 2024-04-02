package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	//exclude this field if empty when encoding (used in GET /api/users to omit password)
	Password string `json:"password,omitempty"`
}

type UserWithToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
