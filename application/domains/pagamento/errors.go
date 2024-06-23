package pagamento

import "errors"

var (
	ErroPedidoIdPagamentoInvalido     = errors.New("o id do pedido é inválido")
	ErroValorPagamentoInvalido        = errors.New("o valor do pagamento deve ser maior que zero")
	ErroStatusPagamentoInvalido       = errors.New("o status do pagamento deve ser: Concluído, Erro ou Pendente")
	ErroEmailClientePagamentoInvalido = errors.New("o e-mail do cliente é inválido")
	ErroTransactionIdObrigatorio      = errors.New("em caso de pagamento bem sucedido é obrigatório informar o transaction id")
)
