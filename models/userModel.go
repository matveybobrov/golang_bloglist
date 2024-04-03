package models

type User struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	//omit this field if empty when encoding (used in GET /api/users to omit password)
	Password string `json:"password,omitempty" db:"password"`
	Username string `json:"username" db:"username"`
}

type UserWithToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
