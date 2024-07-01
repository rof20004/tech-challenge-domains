package cliente

import (
	"strings"

	"github.com/rof20004/tech-challenge-domains/application/utils"

	"github.com/google/uuid"
)

type Cliente struct {
	Id       string `json:"id" bson:"id"`
	Nome     string `json:"nome" bson:"nome"`
	Email    string `json:"email" bson:"email"`
	Cpf      string `json:"cpf" bson:"cpf"`
	Endereco string `json:"endereco" bson:"endereco"`
}

func NovoCliente(nome, email, cpf, endereco string) (Cliente, error) {
	c := Cliente{
		Id:       utils.GenerateUuid(),
		Nome:     nome,
		Email:    email,
		Cpf:      cpf,
		Endereco: endereco,
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

	if strings.TrimSpace(c.Endereco) == "" {
		return ErroEnderecoClienteObrigatorio
	}

	return nil
}
