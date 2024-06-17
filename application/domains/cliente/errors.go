package cliente

import "errors"

var (
	ErroNomeClienteObrigatorio = errors.New("nome do cliente é obrigatório")
	ErroEmailClienteInvalido   = errors.New("e-mail do cliente é inválido")
	ErroCpfClienteInvalido     = errors.New("cpf do cliente é inválido")
)
