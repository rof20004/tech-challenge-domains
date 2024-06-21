package pedido

import "errors"

var (
	ErroIdPedidoInvalido         = errors.New("o id do pedido é inválido")
	ErroStatusPedidoInvalido     = errors.New("o status do pedido deve ser: Recebido, EM_PREPARO, Pronto, Finalizado ou Cancelado")
	ErroValorTotalPedidoInvalido = errors.New("o valor total do pedido deve ser maior que zero")
	ErroClientePedidoInvalido    = errors.New("pedido com dados inválidos do cliente")
	ErroProdutosPedidoInvalidos  = errors.New("existe um ou mais erros nos produtos do pedido")
)
