package cliente

import "errors"

var (
	ErroIdClienteInvalido          = errors.New("o cliente não possui um id válido")
	ErroNomeClienteObrigatorio     = errors.New("nome do cliente é obrigatório")
	ErroEmailClienteInvalido       = errors.New("e-mail do cliente é inválido")
	ErroCpfClienteInvalido         = errors.New("cpf do cliente é inválido")
	ErroEnderecoClienteObrigatorio = errors.New("endereço do cliente é obrigatório")
)
