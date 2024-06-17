package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rof20004/tech-challenge-domains/application/domains/cliente"
)

func TestCliente(t *testing.T) {
	t.Run("deve criar cliente com sucesso", deveCriarClienteComSucesso)
	t.Run("deve validar cliente ao criar", deveValidarClienteAoCriar)
}

func deveCriarClienteComSucesso(t *testing.T) {
	var (
		nome  = "Rodolfo"
		email = "rof20004@gmail.com"
		cpf   = "11122233344"
	)

	c, err := cliente.NovoCliente(nome, email, cpf)
	assert.Nil(t, err)
	assert.NotZero(t, c.Id)
	assert.Equal(t, nome, c.Nome)
	assert.Equal(t, email, c.Email)
	assert.Equal(t, cpf, c.Cpf)
}

func deveValidarClienteAoCriar(t *testing.T) {
	t.Run("validar nome vazio", func(internal *testing.T) {
		var (
			nome  = ""
			email = "rof20004@gmail.com"
			cpf   = "11122233344"
		)

		_, err := cliente.NovoCliente(nome, email, cpf)
		assert.ErrorIs(t, err, cliente.ErroNomeClienteObrigatorio)
	})

	t.Run("validar nome com espaco em branco", func(internal *testing.T) {
		var (
			nome  = "        "
			email = "rof20004@gmail.com"
			cpf   = "11122233344"
		)

		_, err := cliente.NovoCliente(nome, email, cpf)
		assert.ErrorIs(t, err, cliente.ErroNomeClienteObrigatorio)
	})

	t.Run("validar email vazio", func(internal *testing.T) {
		var (
			nome  = "Rodolfo"
			email = ""
			cpf   = "11122233344"
		)

		_, err := cliente.NovoCliente(nome, email, cpf)
		assert.ErrorIs(t, err, cliente.ErroEmailClienteInvalido)
	})

	t.Run("validar email com espaco em branco", func(internal *testing.T) {
		var (
			nome  = "Rodolfo"
			email = "              "
			cpf   = "11122233344"
		)

		_, err := cliente.NovoCliente(nome, email, cpf)
		assert.ErrorIs(t, err, cliente.ErroEmailClienteInvalido)
	})

	t.Run("validar email com formato invalido", func(internal *testing.T) {
		var (
			nome  = "Rodolfo"
			email = "sadasdaidadasduashduashduahdaud"
			cpf   = "11122233344"
		)

		_, err := cliente.NovoCliente(nome, email, cpf)
		assert.ErrorIs(t, err, cliente.ErroEmailClienteInvalido)
	})

	t.Run("validar cpf vazio", func(internal *testing.T) {
		var (
			nome  = "Rodolfo"
			email = "rof20004@gmail.com"
			cpf   = ""
		)

		_, err := cliente.NovoCliente(nome, email, cpf)
		assert.ErrorIs(t, err, cliente.ErroCpfClienteInvalido)
	})

	t.Run("validar cpf com espaco em branco", func(internal *testing.T) {
		var (
			nome  = "Rodolfo"
			email = "rof20004@gmail.com"
			cpf   = "                 "
		)

		_, err := cliente.NovoCliente(nome, email, cpf)
		assert.ErrorIs(t, err, cliente.ErroCpfClienteInvalido)
	})

	t.Run("validar cpf com quantidade menor de caracteres", func(internal *testing.T) {
		var (
			nome  = "Rodolfo"
			email = "rof20004@gmail.com"
			cpf   = "1112223334"
		)

		_, err := cliente.NovoCliente(nome, email, cpf)
		assert.ErrorIs(t, err, cliente.ErroCpfClienteInvalido)
	})

	t.Run("validar cpf com quantidade maior de caracteres", func(internal *testing.T) {
		var (
			nome  = "Rodolfo"
			email = "rof20004@gmail.com"
			cpf   = "1112223334445"
		)

		_, err := cliente.NovoCliente(nome, email, cpf)
		assert.ErrorIs(t, err, cliente.ErroCpfClienteInvalido)
	})
}
