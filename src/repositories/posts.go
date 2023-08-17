package repositories

import (
	"database/sql"
	"fmt"
	"portfolio-api/src/models"
)

type posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *posts {
	return &posts{db}
}

func (repository posts) Create(post models.Post) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO post (name, description, text, user_id) VALUES (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	fmt.Println(post.UserId)

	result, erro := statement.Exec(post.Name, post.Description, post.Text, post.UserId)
	if erro != nil {
		return 0, erro
	}

	id, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(id), nil
}
