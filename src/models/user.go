package models

import (
	"errors"
	"portfolio-api/src/security"
	"strings"

	"github.com/badoux/checkmail"
)

type User struct {
	Id       uint64 `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *User) Prepare(step string) error {
	if erro := u.validar(step); erro != nil {
		return erro
	}

	if erro := u.formatar(step); erro != nil {
		return erro
	}

	return nil
}

func (u *User) validar(step string) error {
	if u.Name == "" {
		return errors.New("O nome não pode estar vazio")
	}

	if u.Email == "" {
		return errors.New("O e-mail não pode estar vazio")
	}

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("O e-mail não atende ao formato xxxxx@email.com")
	}

	if u.Password == "" && step == "create" {
		return errors.New("A senha não pode estar vazio")
	}

	return nil
}

func (u *User) formatar(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)

	if step == "create" {
		passwordHash, erro := security.Hashing(u.Password)
		if erro != nil {
			return erro
		}

		u.Password = string(passwordHash)
	}

	return nil
}
