package unit

import (
	"testing"

	"github.com/rof20004/tech-challenge-domains/application/domains/produto"

	"github.com/stretchr/testify/assert"
)

func TestProduto(t *testing.T) {
	t.Run("deve criar produto com sucesso", deveCriarProdutoComSucesso)
	t.Run("deve validar produto ao criar", deveValidarProdutoAoCriar)
	t.Run("deve atualizar produto com sucesso", deveAtualizarProdutoComSucesso)
	t.Run("deve validar produto ao atualizar", deveValidarProdutoAoAtualizar)
}

func deveCriarProdutoComSucesso(t *testing.T) {
	var (
		nome            = "Misto duplo"
		descricao       = "Sanduiche"
		preco     int64 = 1200
		tipo            = produto.Lanche
	)

	p, err := produto.NovoProduto(nome, descricao, preco, tipo)
	assert.Nil(t, err)
	assert.NotZero(t, p.Id)
	assert.Equal(t, nome, p.Nome)
	assert.Equal(t, descricao, p.Descricao)
	assert.Equal(t, preco, p.Preco)
	assert.Equal(t, tipo, p.Tipo)
}

func deveValidarProdutoAoCriar(t *testing.T) {
	t.Run("validar nome vazio", func(internal *testing.T) {
		var (
			nome            = ""
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.Lanche
		)

		_, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.ErrorIs(internal, err, produto.ErroNomeProdutoObrigatorio)
	})

	t.Run("validar nome com espaco em branco", func(internal *testing.T) {
		var (
			nome            = "               "
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.Lanche
		)

		_, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.ErrorIs(internal, err, produto.ErroNomeProdutoObrigatorio)
	})

	t.Run("validar preco igual a zero", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = 0
			tipo            = produto.Lanche
		)

		_, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.ErrorIs(internal, err, produto.ErroPrecoProdutoInvalido)
	})

	t.Run("validar preco menor zero", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = -100
			tipo            = produto.Lanche
		)

		_, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.ErrorIs(internal, err, produto.ErroPrecoProdutoInvalido)
	})

	t.Run("validar tipo vazio", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.TipoProduto("")
		)

		_, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.ErrorIs(internal, err, produto.ErroTipoProdutoInvalido)
	})

	t.Run("validar tipo com espaco em branco", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.TipoProduto("           ")
		)

		_, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.ErrorIs(internal, err, produto.ErroTipoProdutoInvalido)
	})

	t.Run("validar tipo inexistente", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.TipoProduto("REGIONAL")
		)

		_, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.ErrorIs(internal, err, produto.ErroTipoProdutoInvalido)
	})
}

func deveAtualizarProdutoComSucesso(t *testing.T) {
	var (
		nome            = "Misto duplo"
		descricao       = "Sanduiche"
		preco     int64 = 1200
		tipo            = produto.Lanche
	)

	p, err := produto.NovoProduto(nome, descricao, preco, tipo)
	assert.Nil(t, err)
	assert.NotZero(t, p.Id)
	assert.Equal(t, nome, p.Nome)
	assert.Equal(t, descricao, p.Descricao)
	assert.Equal(t, preco, p.Preco)
	assert.Equal(t, tipo, p.Tipo)
	assert.NotZero(t, p.CriadoEm)
	assert.Zero(t, p.AtualizadoEm)

	err = p.Atualizar("Novo nome", "nova descrição", 1000, produto.Acompanhamento)
	assert.Nil(t, err)
	assert.NotZero(t, p.Id)
	assert.NotEqual(t, nome, p.Nome)
	assert.NotEqual(t, descricao, p.Descricao)
	assert.NotEqual(t, preco, p.Preco)
	assert.NotEqual(t, tipo, p.Tipo)
	assert.NotZero(t, p.AtualizadoEm)
}

func deveValidarProdutoAoAtualizar(t *testing.T) {
	t.Run("validar nome vazio", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.Lanche
		)

		p, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.Nil(t, err)
		assert.NotZero(t, p.Id)
		assert.Equal(t, nome, p.Nome)
		assert.Equal(t, descricao, p.Descricao)
		assert.Equal(t, preco, p.Preco)
		assert.Equal(t, tipo, p.Tipo)

		err = p.Atualizar("", descricao, preco, tipo)
		assert.ErrorIs(internal, err, produto.ErroNomeProdutoObrigatorio)
	})

	t.Run("validar nome com espaco em branco", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.Lanche
		)

		p, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.Nil(t, err)
		assert.NotZero(t, p.Id)
		assert.Equal(t, nome, p.Nome)
		assert.Equal(t, descricao, p.Descricao)
		assert.Equal(t, preco, p.Preco)
		assert.Equal(t, tipo, p.Tipo)

		err = p.Atualizar("            ", descricao, preco, tipo)
		assert.ErrorIs(internal, err, produto.ErroNomeProdutoObrigatorio)
	})

	t.Run("validar preco igual a zero", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.Lanche
		)

		p, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.Nil(t, err)
		assert.NotZero(t, p.Id)
		assert.Equal(t, nome, p.Nome)
		assert.Equal(t, descricao, p.Descricao)
		assert.Equal(t, preco, p.Preco)
		assert.Equal(t, tipo, p.Tipo)

		err = p.Atualizar(nome, descricao, 0, tipo)
		assert.ErrorIs(internal, err, produto.ErroPrecoProdutoInvalido)
	})

	t.Run("validar preco menor que zero", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.Lanche
		)

		p, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.Nil(t, err)
		assert.NotZero(t, p.Id)
		assert.Equal(t, nome, p.Nome)
		assert.Equal(t, descricao, p.Descricao)
		assert.Equal(t, preco, p.Preco)
		assert.Equal(t, tipo, p.Tipo)

		err = p.Atualizar(nome, descricao, -100, tipo)
		assert.ErrorIs(internal, err, produto.ErroPrecoProdutoInvalido)
	})

	t.Run("validar tipo vazio", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.Lanche
		)

		p, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.Nil(t, err)
		assert.NotZero(t, p.Id)
		assert.Equal(t, nome, p.Nome)
		assert.Equal(t, descricao, p.Descricao)
		assert.Equal(t, preco, p.Preco)
		assert.Equal(t, tipo, p.Tipo)

		err = p.Atualizar(nome, descricao, 1200, "")
		assert.ErrorIs(internal, err, produto.ErroTipoProdutoInvalido)
	})

	t.Run("validar tipo com espaco em branco", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.Lanche
		)

		p, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.Nil(t, err)
		assert.NotZero(t, p.Id)
		assert.Equal(t, nome, p.Nome)
		assert.Equal(t, descricao, p.Descricao)
		assert.Equal(t, preco, p.Preco)
		assert.Equal(t, tipo, p.Tipo)

		err = p.Atualizar(nome, descricao, 1200, "                    ")
		assert.ErrorIs(internal, err, produto.ErroTipoProdutoInvalido)
	})

	t.Run("validar tipo inexistente", func(internal *testing.T) {
		var (
			nome            = "Misto duplo"
			descricao       = "Sanduiche"
			preco     int64 = 1200
			tipo            = produto.Lanche
		)

		p, err := produto.NovoProduto(nome, descricao, preco, tipo)
		assert.Nil(t, err)
		assert.NotZero(t, p.Id)
		assert.Equal(t, nome, p.Nome)
		assert.Equal(t, descricao, p.Descricao)
		assert.Equal(t, preco, p.Preco)
		assert.Equal(t, tipo, p.Tipo)

		err = p.Atualizar(nome, descricao, 1200, "TESTE")
		assert.ErrorIs(internal, err, produto.ErroTipoProdutoInvalido)
	})
}
