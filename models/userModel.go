package models

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	//exclude this field if empty when encoding (used in GET /api/users to omit password)
	Password string `json:"password,omitempty"`
	Username string `json:"username"`
}

type UserWithToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
