package db

import "bloglist/models"

type User = models.User

func GetUserById(id int) (User, error) {
	user := User{}
	row := DB.QueryRow("SELECT * FROM users WHERE id=$1", id)
	err := row.Scan(&user.Id, &user.Username, &user.Name, &user.Password)
	return user, err
}

func GetUserByUsername(username string) (User, error) {
	user := User{}
	row := DB.QueryRow("SELECT * FROM users WHERE username=$1", username)
	err := row.Scan(&user.Id, &user.Username, &user.Name, &user.Password)
	return user, err
}

func GetAllUsers() ([]User, error) {
	users := []User{}
	rows, err := DB.Query("SELECT id, username, name FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func InsertUser(user User) (User, error) {
	savedUser := User{}
	row := DB.QueryRow("INSERT INTO users (username, name, password) VALUES ($1, $2, $3) RETURNING id, username, name", user.Username, user.Name, user.Password)
	err := row.Scan(&savedUser.Id, &savedUser.Username, &savedUser.Name)
	return savedUser, err
}
