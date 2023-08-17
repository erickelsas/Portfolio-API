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
	"portfolio-api/src/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	bd, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repository := repositories.NewUserRepository(bd)
	userDB, erro := repository.SearchUserByEmail(user.Email)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(userDB.Password, user.Password); erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := authorization.TokenCreate(userDB.Id)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	w.Write([]byte(token))
}
