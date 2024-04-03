package db

import "bloglist/models"

type User = models.User

func GetUserById(id int) (User, error) {
	user := User{}

	row := DB.QueryRow(`
		SELECT
			id, name, username, password
    FROM users
    WHERE
			id = $1
	`, id)

	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Password,
	)
	return user, err
}

func GetUserByUsername(username string) (User, error) {
	user := User{}

	row := DB.QueryRow(`
    SELECT 
			id, name, username, password
    FROM users
    WHERE
			username = $1
	`, username)

	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Password,
	)
	return user, err
}

func GetAllUsers() ([]User, error) {
	users := []User{}

	rows, err := DB.Query(`
    SELECT
    	id, name, username
    FROM users
	`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := User{}
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Username,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func InsertUser(user User) (User, error) {
	savedUser := User{}

	row := DB.QueryRow(`
		INSERT INTO users
			(name, username, password)
		VALUES
			($1, $2, $3)
		RETURNING
			id, name, username, password
	`, user.Name, user.Username, user.Password)

	err := row.Scan(
		&savedUser.Id,
		&savedUser.Name,
		&savedUser.Username,
		&savedUser.Password,
	)
	return savedUser, err
}
