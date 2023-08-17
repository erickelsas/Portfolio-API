package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	Id          uint64    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Text        string    `json:"text, omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UserId      uint64    `json:"userId,omitempty"`
}

func (p *Post) Prepare() error {
	if erro := p.validar(); erro != nil {
		return erro
	}

	p.formatar()

	return nil
}

func (p *Post) validar() error {
	if p.Name == "" {
		return errors.New("O título da postagem não pode ser vazio")
	}

	if p.Description == "" {
		return errors.New("A descrição da postagem não pode ser vazia")
	}

	return nil
}

func (p *Post) formatar() {
	p.Name = strings.TrimSpace(p.Name)
	p.Description = strings.TrimSpace(p.Description)
	p.Text = strings.TrimSpace(p.Text)
}
