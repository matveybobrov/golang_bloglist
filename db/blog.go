package db

import (
	"bloglist/models"
)

type Blog = models.Blog
type BlogWithUser = models.BlogWithUser

func GetAllBlogs() ([]Blog, error) {
	blogs := []Blog{}

	rows, err := DB.Query(`
		SELECT  *
    FROM blogs
	`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		blog := Blog{}
		// Scan should match the db schema fields order
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

func GetAllBlogsWithUsers() ([]BlogWithUser, error) {
	blogs := []BlogWithUser{}

	rows, err := DB.Query(`
		SELECT  *
    FROM blogs
    JOIN users
      ON blogs.user_id = users.id
	`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		blog := BlogWithUser{}
		err := rows.Scan(
			&blog.Id,
			&blog.Title,
			&blog.Author,
			&blog.Url,
			&blog.Likes,
			&blog.User_id,

			&blog.User.Id,
			&blog.User.Name,
			&blog.User.Password,
			&blog.User.Username,
		)
		if err != nil {
			return nil, err
		}
		blog.User.Password = ""
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func GetBlogById(id int) (Blog, error) {
	blog := Blog{}

	row := DB.QueryRow(`
		SELECT  *
    FROM blogs
    WHERE id = $1 
	`, id)

	err := row.Scan(
		&blog.Id,
		&blog.Author,
		&blog.Url,
		&blog.Title,
		&blog.Likes,
		&blog.User_id,
	)
	if err != nil {
		return blog, err
	}
	return blog, nil
}

func GetBlogWithUserById(id int) (BlogWithUser, error) {
	blog := BlogWithUser{}

	row := DB.QueryRow(`
		SELECT  *
    FROM blogs
    JOIN users
      ON blogs.user_id = users.id
    WHERE blogs.id = $1
	`, id)

	err := row.Scan(
		&blog.Id,
		&blog.Author,
		&blog.Url,
		&blog.Title,
		&blog.Likes,
		&blog.User_id,

		&blog.User.Id,
		&blog.User.Name,
		&blog.User.Password,
		&blog.User.Username,
	)
	blog.User.Password = ""
	if err != nil {
		return blog, err
	}
	return blog, nil
}

func InsertBlog(blog Blog) (Blog, error) {
	savedBlog := Blog{}

	row := DB.QueryRow(`
		INSERT INTO blogs
		(title, author, url, user_id)
		VALUES
		($1, $2, $3, $4)
		RETURNING *
	`, blog.Title, blog.Author, blog.Url, blog.User_id)

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
	_, err := DB.Exec(`
		DELETE
    FROM blogs
    WHERE id = $1
	`, id)
	return err
}

func UpdateBlogById(blog Blog, id int) (Blog, error) {
	updatedBlog := Blog{}

	row := DB.QueryRow(`
		UPDATE blogs
    SET title = $2, author = $3, url = $4, likes = $5
    WHERE id = $1 RETURNING *
	`, id, blog.Title, blog.Author, blog.Url, blog.Likes)

	err := row.Scan(
		&updatedBlog.Id,
		&updatedBlog.Title,
		&updatedBlog.Author,
		&updatedBlog.Url,
		&updatedBlog.Likes,
		&updatedBlog.User_id,
	)
	if err != nil {
		return updatedBlog, err
	}
	return updatedBlog, nil
}
