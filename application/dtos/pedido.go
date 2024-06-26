package dtos

import (
	"github.com/rof20004/tech-challenge-domains/application/domains/cliente"
	"github.com/rof20004/tech-challenge-domains/application/domains/pedido"
	"github.com/rof20004/tech-challenge-domains/application/domains/produto"
)

type CriarPedidoDto struct {
	Cliente  cliente.Cliente   `json:"cliente"`
	Produtos []produto.Produto `json:"produtos"`
	Cartao   CartaoDto         `json:"cartao"`
}

type AtualizarPedidoDto struct {
	PedidoId string              `json:"pedidoId"`
	Status   pedido.StatusPedido `json:"status"`
}
