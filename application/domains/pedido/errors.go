package pedido

import "errors"

var (
	ErroStatusPedidoInvalido     = errors.New("o status do pedido deve ser: RECEBIDO, EM_PREPARO, PRONTO, FINALIZADO ou CANCELADO")
	ErroValorTotalPedidoInvalido = errors.New("o valor total do pedido deve ser maior que zero")
	ErroClientePedidoInvalido    = errors.New("pedido com dados inválidos do cliente")
)
