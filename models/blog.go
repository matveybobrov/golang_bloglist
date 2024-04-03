package models

type Blog struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Url     string `json:"url"`
	Likes   int    `json:"likes"`
	User_id int    `json:"user_id"`
}
