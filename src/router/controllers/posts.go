package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"portfolio-api/src/authorization"
	"portfolio-api/src/db"
	"portfolio-api/src/models"
	"portfolio-api/src/repositories"
	"portfolio-api/src/responses"
)

// CreatePost cria um novo postagem
func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Post
	if erro = json.Unmarshal(body, &post); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = post.Prepare(); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	post.UserId, erro = authorization.GetUserId(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
	}

	bd, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
	}
	defer bd.Close()

	repository := repositories.NewPostRepository(bd)
	post.Id, erro = repository.Create(post)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

// SearchPosts busca todos os postagens do banco
func SearchPosts(w http.ResponseWriter, r *http.Request) {

}

// SearchPostById busca um postagem pelo ID
func SearchPostById(w http.ResponseWriter, r *http.Request) {

}

// UpdatePost atualiza os dados de um postagem
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

// DeletePost apaga um postagem pelo ID
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
