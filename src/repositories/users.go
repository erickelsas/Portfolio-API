package repositories

import (
	"database/sql"
	"fmt"
	"portfolio-api/src/models"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO user (name, email, password) VALUES (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	id, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(id), nil
}

func (repository users) SearchUserByName(name string) ([]models.User, error) {
	name = fmt.Sprintf("%%%s%%", name)
	lines, erro := repository.db.Query("SELECT id, name, email FROM user WHERE name LIKE ?", name)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if erro = lines.Scan(&user.Id, &user.Name, &user.Email); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository users) SearchUserById(id uint64) (models.User, error) {
	line, erro := repository.db.Query("SELECT id, name, email FROM user WHERE id = ?", id)
	if erro != nil {
		return models.User{}, erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.Id, &user.Name, &user.Email); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repository users) SearchUserByEmail(email string) (models.User, error) {
	line, erro := repository.db.Query("SELECT id, password FROM user WHERE email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.Id, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repository users) UpdateUser(id uint64, user models.User) error {
	statement, erro := repository.db.Prepare("UPDATE user SET name = ?, email = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(user.Name, user.Email, id); erro != nil {
		return erro
	}

	return nil
}

func (repository users) DeleteUser(id uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM user WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}
