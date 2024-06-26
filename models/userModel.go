package models

// those db tags are coming from sqlx
type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Name     string `json:"name" db:"name"`
	// omit this field if empty when encoding
	// used for hiding user password in response
	Password string `json:"password,omitempty" db:"password"`
}

type UserWithToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
