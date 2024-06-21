package unit

import (
	"testing"

	"github.com/rof20004/tech-challenge-domains/application/domains/pagamento"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPagamento(t *testing.T) {
	t.Run("deve criar pagamento com sucesso", deveCriarPagamentoComSucesso)
	t.Run("deve validar pagamento ao criar", deveValidarPagamentoAoCriar)
	t.Run("deve atualizar pagamento com sucesso", deveAtualizarPagamentoComSucesso)
	t.Run("deve validar pagamento ao atualizar", deveValidarPagamentoAoAtualizar)
}

func deveCriarPagamentoComSucesso(t *testing.T) {
	var (
		pedidoId       = uuid.NewString()
		valor    int64 = 2000
	)

	p, err := pagamento.NovoPagamento(pedidoId, valor)
	assert.Nil(t, err)
	assert.NotZero(t, p.Id)
	assert.Equal(t, pedidoId, p.PedidoId)
	assert.Equal(t, valor, p.Valor)
	assert.Equal(t, pagamento.Pendente, p.Status)
	assert.NotZero(t, p.CriadoEm)
}

func deveValidarPagamentoAoCriar(t *testing.T) {
	t.Run("validar pedido id vazio", func(internal *testing.T) {
		var (
			pedidoId       = ""
			valor    int64 = 2000
		)

		_, err := pagamento.NovoPagamento(pedidoId, valor)
		assert.ErrorIs(internal, err, pagamento.ErroPedidoIdPagamentoInvalido)
	})

	t.Run("validar pedido id com espaco em branco", func(internal *testing.T) {
		var (
			pedidoId       = "         "
			valor    int64 = 2000
		)

		_, err := pagamento.NovoPagamento(pedidoId, valor)
		assert.ErrorIs(internal, err, pagamento.ErroPedidoIdPagamentoInvalido)
	})

	t.Run("validar pedido id invalido", func(internal *testing.T) {
		var (
			pedidoId       = "344j4njn234j"
			valor    int64 = 2000
		)

		_, err := pagamento.NovoPagamento(pedidoId, valor)
		assert.ErrorIs(internal, err, pagamento.ErroPedidoIdPagamentoInvalido)
	})

	t.Run("validar valor igual a zero", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 0
		)

		_, err := pagamento.NovoPagamento(pedidoId, valor)
		assert.ErrorIs(internal, err, pagamento.ErroValorPagamentoInvalido)
	})

	t.Run("validar valor igual menor que zero", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = -1000
		)

		_, err := pagamento.NovoPagamento(pedidoId, valor)
		assert.ErrorIs(internal, err, pagamento.ErroValorPagamentoInvalido)
	})
}

func deveAtualizarPagamentoComSucesso(t *testing.T) {
	var (
		pedidoId       = uuid.NewString()
		valor    int64 = 2000
	)

	p, err := pagamento.NovoPagamento(pedidoId, valor)
	assert.Nil(t, err)
	assert.NotZero(t, p.Id)
	assert.Equal(t, pedidoId, p.PedidoId)
	assert.Equal(t, valor, p.Valor)
	assert.Equal(t, pagamento.Pendente, p.Status)
	assert.NotZero(t, p.CriadoEm)

	err = p.Atualizar(pagamento.Concluido)
	assert.Nil(t, err)
	assert.Equal(t, pagamento.Concluido, p.Status)
	assert.NotZero(t, p.AtualizadoEm)
}

func deveValidarPagamentoAoAtualizar(t *testing.T) {
	t.Run("validar status vazio", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 2000
		)

		p, err := pagamento.NovoPagamento(pedidoId, valor)
		assert.Nil(internal, err)

		err = p.Atualizar("")
		assert.ErrorIs(internal, err, pagamento.ErroStatusPagamentoInvalido)
	})

	t.Run("validar status com espaco em branco", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 2000
		)

		p, err := pagamento.NovoPagamento(pedidoId, valor)
		assert.Nil(internal, err)

		err = p.Atualizar("           ")
		assert.ErrorIs(internal, err, pagamento.ErroStatusPagamentoInvalido)
	})

	t.Run("validar status invalido", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 2000
		)

		p, err := pagamento.NovoPagamento(pedidoId, valor)
		assert.Nil(internal, err)

		err = p.Atualizar("TESTE")
		assert.ErrorIs(internal, err, pagamento.ErroStatusPagamentoInvalido)
	})
}
