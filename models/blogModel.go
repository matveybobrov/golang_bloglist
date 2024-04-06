package models

// struct tags such as json and db help to parse data correctly
type Blog struct {
	Id     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
	Url    string `json:"url" db:"url"`
	Likes  int    `json:"likes" db:"likes"`
	// exclude this field completely from json because of user population
	// it will be completely ommited on marshall and unmarshall
	UserId int `json:"-" db:"user_id"`
}

// all fields from Blog will be embedded into this type
type BlogWithUser struct {
	Blog
	User User `json:"user"`
}
