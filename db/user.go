package db

import "bloglist/models"

type User = models.User

func InsertUser(user User) (User, error) {
	savedUser := User{}
	row := DB.QueryRow("INSERT INTO users (username, name, password) VALUES ($1, $2, $3) RETURNING *", user.Username, user.Name, user.Password)
	err := row.Scan(&savedUser.Id, &savedUser.Username, &savedUser.Name, &savedUser.Password)
	if err != nil {
		return savedUser, err
	}
	return savedUser, nil
}
