package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"portfolio-api/src/db"
	"portfolio-api/src/models"
	"portfolio-api/src/repositories"
	"portfolio-api/src/responses"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser cria um novo usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(body, &user); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("create"); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
	}
	defer bd.Close()

	repository := repositories.NewUserRepository(bd)
	user.Id, erro = repository.Create(user)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// SearchUsers busca todos os usuários do banco
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Query().Get("name"))

	bd, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repository := repositories.NewUserRepository(bd)
	users, erro := repository.SearchUserByName(name)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// SearchUserById busca um usuário pelo ID
func SearchUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	bd, erro = db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repository := repositories.NewUserRepository(bd)
	user, erro := repository.SearchUserById(userId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser atualiza os dados de um usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(body, &user); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("update"); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repository := repositories.NewUserRepository(bd)
	if erro = repository.UpdateUser(userId, user); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser apaga um usuário pelo ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repository := repositories.NewUserRepository(bd)
	if erro = repository.DeleteUser(userId); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
