package pagamento

import "errors"

var (
	ErroPedidoIdPagamentoInvalido = errors.New("o id do pedido é inválido")
	ErroValorPagamentoInvalido    = errors.New("o valor do pagamento deve ser maior que zero")
	ErroStatusPagamentoInvalido   = errors.New("o status do pagamento deve ser: Concluído, Erro ou Pendente")
)
