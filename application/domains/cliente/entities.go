package cliente

import (
	"strings"

	"github.com/rof20004/tech-challenge-domains/application/utils"

	"github.com/google/uuid"
)

type Cliente struct {
	Id    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
}

func NovoCliente(nome, email, cpf string) (Cliente, error) {
	c := Cliente{
		Id:    utils.GenerateUuid(),
		Nome:  nome,
		Email: email,
		Cpf:   cpf,
	}

	return c, c.Validar()
}

func (c Cliente) Validar() error {
	if _, err := uuid.Parse(c.Id); err != nil {
		return ErroIdClienteInvalido
	}

	if strings.TrimSpace(c.Nome) == "" {
		return ErroNomeClienteObrigatorio
	}

	if !utils.IsEmailValido(c.Email) {
		return ErroEmailClienteInvalido
	}

	if strings.TrimSpace(c.Cpf) == "" || len(c.Cpf) != 11 {
		return ErroCpfClienteInvalido
	}

	return nil
}
