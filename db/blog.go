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
		err := rows.Scan(
			&blog.Id,
			&blog.Title,
			&blog.Author,
			&blog.Url,
			&blog.Likes,
			&blog.User_id,
		)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func GetBlogById(id int) (Blog, error) {
	blog := Blog{}
	row := DB.QueryRow("SELECT * FROM blogs WHERE id=$1", id)
	err := row.Scan(
		&blog.Id,
		&blog.Author,
		&blog.Url,
		&blog.Title,
		&blog.Likes,
		&blog.User_id,
	)
	if err == sql.ErrNoRows {
		return blog, err
	}
	if err != nil {
		return blog, err
	}
	return blog, nil
}

func InsertBlog(blog Blog) (Blog, error) {
	savedBlog := Blog{}
	row := DB.QueryRow("INSERT INTO blogs (title, author, url, user_id) VALUES ($1, $2, $3, $4) RETURNING *",
		blog.Title,
		blog.Author,
		blog.Url,
		blog.User_id,
	)
	err := row.Scan(
		&savedBlog.Id,
		&savedBlog.Title,
		&savedBlog.Author,
		&savedBlog.Url,
		&savedBlog.Likes,
		&savedBlog.User_id,
	)
	if err != nil {
		return savedBlog, err
	}
	return savedBlog, nil
}

func DeleteBlogById(id int) error {
	_, err := DB.Exec("DELETE FROM blogs WHERE id=$1", id)
	return err
}

func UpdateBlogById(blog Blog, id int) (Blog, error) {
	updatedBlog := Blog{}
	row := DB.QueryRow("UPDATE blogs SET title=$2, author=$3, url=$4, likes=$5 WHERE id=$1 RETURNING *", blog.Id, blog.Title, blog.Author, blog.Url, blog.Likes)
	err := row.Scan(&updatedBlog.Id, &updatedBlog.Title, &updatedBlog.Author, &updatedBlog.Url, &updatedBlog.Likes)
	if err != nil {
		return updatedBlog, err
	}
	return updatedBlog, nil
}
