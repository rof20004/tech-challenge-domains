package unit

import (
	"testing"

	"github.com/rof20004/tech-challenge-domains/application/domains/cliente"
	"github.com/rof20004/tech-challenge-domains/application/domains/pedido"
	"github.com/rof20004/tech-challenge-domains/application/domains/produto"

	"github.com/stretchr/testify/assert"
)

func TestPedido(t *testing.T) {
	t.Run("deve criar pedido com sucesso", deveCriarPedidoComSucesso)
	t.Run("deve validar pedido ao criar", deveValidarPedidoAoCriar)
	t.Run("deve atualizar pedido com sucesso", deveAtualizarPedidoComSucesso)
	t.Run("deve validar pedido ao atualizar", deveValidarPedidoAoAtualizar)
}

func deveCriarPedidoComSucesso(t *testing.T) {
	cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
	assert.Nil(t, err)

	produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
	assert.Nil(t, err)

	produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)

	produtos := []produto.Produto{produto1, produto2}

	p, err := pedido.NovoPedido(cl, produtos)
	assert.Nil(t, err)
	assert.NotZero(t, p.Id)
	assert.Equal(t, cl, p.Cliente)
	assert.Equal(t, produtos, p.Produtos)
	assert.Equal(t, int64(2100), p.ValorTotal)
	assert.NotZero(t, p.CriadoEm)
	assert.Zero(t, p.AtualizadoEm)
}

func deveValidarPedidoAoCriar(t *testing.T) {
	t.Run("validar id do cliente vazio", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Id = ""
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar id do cliente com espaco em branco", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Id = "           "
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar id do cliente invalido", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Id = "132131233912i3193i1391i"
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar nome do cliente vazio", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Nome = ""
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar nome do cliente com espaco em branco", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Nome = "          "
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar email do cliente vazio", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Email = ""
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar email do cliente com espaco em branco", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Email = "           "
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar email do cliente com formato invalido", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Email = "31231231231312312"
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar cpf do cliente vazio", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Cpf = ""
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar cpf do cliente com espaco em branco", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Cpf = "           "
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar cpf do cliente com quantidade maior de caracteres", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Cpf = "1112223334"
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar cpf do cliente com quantidade maior de caracteres", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Cpf = "111222333445"
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar endereco do cliente vazio", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Endereco = ""
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar endereco do cliente com espaco em branco", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produtos := []produto.Produto{produto1, produto2}

		cl.Endereco = "             "
		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroClientePedidoInvalido)
	})

	t.Run("validar produtos com um produto de id vazio", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produto1.Id = ""

		_, err = pedido.NovoPedido(cl, []produto.Produto{produto1, produto2})
		assert.ErrorIs(internal, err, pedido.ErroProdutosPedidoInvalidos)
	})

	t.Run("validar produtos com um produto de preco igual a zero", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
		assert.Nil(internal, err)

		produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)
		assert.Nil(internal, err)

		produto2.Preco = 0

		_, err = pedido.NovoPedido(cl, []produto.Produto{produto1, produto2})
		assert.ErrorIs(internal, err, pedido.ErroProdutosPedidoInvalidos)
	})

	t.Run("validar valor total passando produtos vazio", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		var produtos []produto.Produto

		_, err = pedido.NovoPedido(cl, produtos)
		assert.ErrorIs(internal, err, pedido.ErroValorTotalPedidoInvalido)
	})

	t.Run("validar valor total passando produtos nil", func(internal *testing.T) {
		cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
		assert.Nil(internal, err)

		_, err = pedido.NovoPedido(cl, nil)
		assert.ErrorIs(internal, err, pedido.ErroValorTotalPedidoInvalido)
	})
}

func deveAtualizarPedidoComSucesso(t *testing.T) {
	cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
	assert.Nil(t, err)

	produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
	assert.Nil(t, err)

	produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)

	produtos := []produto.Produto{produto1, produto2}

	p, err := pedido.NovoPedido(cl, produtos)
	assert.Nil(t, err)
	assert.NotZero(t, p.Id)
	assert.Equal(t, cl, p.Cliente)
	assert.Equal(t, produtos, p.Produtos)
	assert.Equal(t, int64(2100), p.ValorTotal)
	assert.NotZero(t, p.CriadoEm)
	assert.Zero(t, p.AtualizadoEm)

	err = p.Atualizar(pedido.CANCELADO)
	assert.Nil(t, err)
	assert.Equal(t, pedido.CANCELADO, p.Status)
}

func deveValidarPedidoAoAtualizar(t *testing.T) {
	cl, err := cliente.NovoCliente("Rodolfo", "rof20004@gmail.com", "11122233344", "Rua A")
	assert.Nil(t, err)

	produto1, err := produto.NovoProduto("Hot dog", "", 900, produto.LANCHE)
	assert.Nil(t, err)

	produto2, err := produto.NovoProduto("Misto duplo", "", 1200, produto.LANCHE)

	produtos := []produto.Produto{produto1, produto2}

	p, err := pedido.NovoPedido(cl, produtos)
	assert.Nil(t, err)
	assert.NotZero(t, p.Id)
	assert.Equal(t, cl, p.Cliente)
	assert.Equal(t, produtos, p.Produtos)
	assert.Equal(t, int64(2100), p.ValorTotal)
	assert.NotZero(t, p.CriadoEm)
	assert.Zero(t, p.AtualizadoEm)

	t.Run("validar status vazio", func(internal *testing.T) {
		err := p.Atualizar("")
		assert.ErrorIs(internal, err, pedido.ErroStatusPedidoInvalido)
	})

	t.Run("validar status com espacos em branco", func(internal *testing.T) {
		err := p.Atualizar("             ")
		assert.ErrorIs(internal, err, pedido.ErroStatusPedidoInvalido)
	})

	t.Run("validar status inexistente", func(internal *testing.T) {
		err := p.Atualizar("TESTE")
		assert.ErrorIs(internal, err, pedido.ErroStatusPedidoInvalido)
	})

	t.Run("validar id vazio", func(internal *testing.T) {
		p.Id = ""
		err := p.Atualizar(pedido.FINALIZADO)
		assert.ErrorIs(internal, err, pedido.ErroIdPedidoInvalido)
	})

	t.Run("validar id com espaco em branco", func(internal *testing.T) {
		p.Id = "           "
		err := p.Atualizar(pedido.FINALIZADO)
		assert.ErrorIs(internal, err, pedido.ErroIdPedidoInvalido)
	})

	t.Run("validar id invalido", func(internal *testing.T) {
		p.Id = "1231233123123"
		err := p.Atualizar(pedido.FINALIZADO)
		assert.ErrorIs(internal, err, pedido.ErroIdPedidoInvalido)
	})
}
