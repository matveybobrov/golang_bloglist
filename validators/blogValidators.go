package validators

import "bloglist/models"

type Blog = models.Blog

func ValidateBlog(blog Blog) (errorMessage, isValid) {
	if blog.Title == "" {
		return "Title must be provided", false
	}
	if blog.Author == "" {
		return "Author must be provided", false
	}
	if blog.Url == "" {
		return "Url must be provided", false
	}
	return "", true
}
