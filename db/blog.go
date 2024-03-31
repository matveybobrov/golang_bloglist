package db

import (
	"bloglist/models"
	"database/sql"
)

type Blog = models.Blog

func GetAllBlogs() ([]Blog, error) {
	blogs := []Blog{}
	rows, err := DB.Query("SELECT * FROM blogs")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		blog := Blog{}
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Author, &blog.Url, &blog.Likes)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func GetOneBlog(id int) (Blog, error) {
	blog := Blog{}
	row := DB.QueryRow("SELECT * FROM blogs WHERE id=$1", id)
	err := row.Scan(&blog.Id, &blog.Author, &blog.Url, &blog.Title, &blog.Likes)
	if err == sql.ErrNoRows {
		return blog, err
	}
	if err != nil {
		return blog, err
	}
	return blog, nil
}

func CreateOneBlog(blog Blog) (Blog, error) {
	savedBlog := Blog{}
	row := DB.QueryRow("INSERT INTO blogs (title, author, url) VALUES ($1, $2, $3) RETURNING *", blog.Title, blog.Author, blog.Url)
	err := row.Scan(&savedBlog.Id, &savedBlog.Title, &savedBlog.Author, &savedBlog.Url, &savedBlog.Likes)
	if err != nil {
		return savedBlog, err
	}
	return savedBlog, nil
}

func DeleteOneBlog(id int) error {
	_, err := DB.Exec("DELETE FROM blogs WHERE id=$1", id)
	return err
}

func UpdateOneBlog(blog Blog, id int) (Blog, error) {
	updatedBlog := Blog{}
	row := DB.QueryRow("UPDATE blogs SET title=$1, author=$2, url=$3 RETURNING *", blog.Title, blog.Author, blog.Url)
	err := row.Scan(&updatedBlog.Id, &updatedBlog.Title, &updatedBlog.Author, &updatedBlog.Url, &updatedBlog.Likes)
	if err != nil {
		return updatedBlog, err
	}
	return updatedBlog, nil
}
