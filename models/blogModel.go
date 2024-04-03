package models

type Blog struct {
	Id     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
	Url    string `json:"url" db:"url"`
	Likes  int    `json:"likes" db:"likes"`
	// exclude this field completely from json because of user population
	User_id int `json:"-" db:"user_id"`
}

type BlogWithUser struct {
	Blog
	User User `json:"user"`
}
