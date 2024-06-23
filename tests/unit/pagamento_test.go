package unit

import (
	"errors"
	"testing"

	"github.com/rof20004/tech-challenge-domains/application/domains/pagamento"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPagamento(t *testing.T) {
	t.Run("deve criar pagamento com sucesso", deveCriarPagamentoComSucesso)
	t.Run("deve validar pagamento ao criar", deveValidarPagamentoAoCriar)
	t.Run("deve atualizar pagamento com sucesso", deveAtualizarPagamentoComSucesso)
	t.Run("deve atualizar pagamento com motivo erro", deveAtualizarPagamentoComMotivoErro)
	t.Run("deve validar pagamento ao atualizar", deveValidarPagamentoAoAtualizar)
}

func deveCriarPagamentoComSucesso(t *testing.T) {
	var (
		pedidoId       = uuid.NewString()
		valor    int64 = 2000
		email          = "rof20004@gmail.com"
	)

	p, err := pagamento.NovoPagamento(pedidoId, email, valor)
	assert.Nil(t, err)
	assert.NotZero(t, p.Id)
	assert.Equal(t, pedidoId, p.PedidoId)
	assert.Equal(t, email, p.EmailCliente)
	assert.Equal(t, valor, p.Valor)
	assert.Equal(t, pagamento.Pendente, p.Status)
	assert.NotZero(t, p.CriadoEm)
}

func deveValidarPagamentoAoCriar(t *testing.T) {
	t.Run("validar pedido id vazio", func(internal *testing.T) {
		var (
			pedidoId       = ""
			valor    int64 = 2000
			email          = "rof20004@gmail.com"
		)

		_, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.ErrorIs(internal, err, pagamento.ErroPedidoIdPagamentoInvalido)
	})

	t.Run("validar pedido id com espaco em branco", func(internal *testing.T) {
		var (
			pedidoId       = "         "
			valor    int64 = 2000
			email          = "rof20004@gmail.com"
		)

		_, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.ErrorIs(internal, err, pagamento.ErroPedidoIdPagamentoInvalido)
	})

	t.Run("validar pedido id invalido", func(internal *testing.T) {
		var (
			pedidoId       = "344j4njn234j"
			valor    int64 = 2000
			email          = "rof20004@gmail.com"
		)

		_, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.ErrorIs(internal, err, pagamento.ErroPedidoIdPagamentoInvalido)
	})

	t.Run("validar valor igual a zero", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 0
			email          = "rof20004@gmail.com"
		)

		_, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.ErrorIs(internal, err, pagamento.ErroValorPagamentoInvalido)
	})

	t.Run("validar valor igual menor que zero", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = -1000
			email          = "rof20004@gmail.com"
		)

		_, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.ErrorIs(internal, err, pagamento.ErroValorPagamentoInvalido)
	})

	t.Run("validar email cliente vazio", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 2000
			email          = ""
		)

		_, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.ErrorIs(internal, err, pagamento.ErroEmailClientePagamentoInvalido)
	})

	t.Run("validar email cliente com espaco em branco", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 2000
			email          = "          "
		)

		_, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.ErrorIs(internal, err, pagamento.ErroEmailClientePagamentoInvalido)
	})

	t.Run("validar email cliente invalido", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 2000
			email          = "1i2j31i23j2i1j@tes"
		)

		_, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.ErrorIs(internal, err, pagamento.ErroEmailClientePagamentoInvalido)
	})
}

func deveAtualizarPagamentoComSucesso(t *testing.T) {
	var (
		pedidoId       = uuid.NewString()
		valor    int64 = 2000
		email          = "rof20004@gmail.com"
	)

	p, err := pagamento.NovoPagamento(pedidoId, email, valor)
	assert.Nil(t, err)
	assert.NotZero(t, p.Id)
	assert.Equal(t, pedidoId, p.PedidoId)
	assert.Equal(t, email, p.EmailCliente)
	assert.Equal(t, valor, p.Valor)
	assert.Equal(t, pagamento.Pendente, p.Status)
	assert.NotZero(t, p.CriadoEm)

	err = p.Atualizar(uuid.NewString(), nil, pagamento.Concluido)
	assert.Nil(t, err)
	assert.Equal(t, pagamento.Concluido, p.Status)
	assert.NotZero(t, p.TransactionId)
	assert.Zero(t, p.MotivoErro)
	assert.NotZero(t, p.AtualizadoEm)
}

func deveAtualizarPagamentoComMotivoErro(t *testing.T) {
	var (
		pedidoId       = uuid.NewString()
		valor    int64 = 2000
		email          = "rof20004@gmail.com"
	)

	p, err := pagamento.NovoPagamento(pedidoId, email, valor)
	assert.Nil(t, err)
	assert.NotZero(t, p.Id)
	assert.Equal(t, pedidoId, p.PedidoId)
	assert.Equal(t, email, p.EmailCliente)
	assert.Equal(t, valor, p.Valor)
	assert.Equal(t, pagamento.Pendente, p.Status)
	assert.NotZero(t, p.CriadoEm)

	err = p.Atualizar("", errors.New("erro qualquer"), pagamento.Erro)
	assert.Nil(t, err)
	assert.Equal(t, pagamento.Erro, p.Status)
	assert.Zero(t, p.TransactionId)
	assert.NotZero(t, p.MotivoErro)
	assert.NotZero(t, p.AtualizadoEm)
}

func deveValidarPagamentoAoAtualizar(t *testing.T) {
	t.Run("validar status vazio", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 2000
			email          = "rof20004@gmail.com"
		)

		p, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.Nil(internal, err)

		err = p.Atualizar(uuid.NewString(), nil, "")
		assert.ErrorIs(internal, err, pagamento.ErroStatusPagamentoInvalido)
	})

	t.Run("validar status com espaco em branco", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 2000
			email          = "rof20004@gmail.com"
		)

		p, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.Nil(internal, err)

		err = p.Atualizar(uuid.NewString(), nil, "             ")
		assert.ErrorIs(internal, err, pagamento.ErroStatusPagamentoInvalido)
	})

	t.Run("validar status invalido", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 2000
			email          = "rof20004@gmail.com"
		)

		p, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.Nil(internal, err)

		err = p.Atualizar(uuid.NewString(), nil, "TESTE")
		assert.ErrorIs(internal, err, pagamento.ErroStatusPagamentoInvalido)
	})

	t.Run("validar transaction id vazio e erro nao preenchido", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 2000
			email          = "rof20004@gmail.com"
		)

		p, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.Nil(internal, err)

		err = p.Atualizar("", nil, pagamento.Concluido)
		assert.ErrorIs(internal, err, pagamento.ErroTransactionIdObrigatorio)
	})

	t.Run("validar transaction id com espaco em branco e erro nao preenchido", func(internal *testing.T) {
		var (
			pedidoId       = uuid.NewString()
			valor    int64 = 2000
			email          = "rof20004@gmail.com"
		)

		p, err := pagamento.NovoPagamento(pedidoId, email, valor)
		assert.Nil(internal, err)

		err = p.Atualizar("        ", nil, pagamento.Concluido)
		assert.ErrorIs(internal, err, pagamento.ErroTransactionIdObrigatorio)
	})
}
