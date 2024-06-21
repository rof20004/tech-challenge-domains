package dtos

import (
	"time"

	"github.com/rof20004/tech-challenge-domains/application/domains/pagamento"
)

type CartaoDto struct {
	Bandeira    string `json:"bandeira"`
	Numero      string `json:"numero"`
	MesValidade string `json:"mesValidade"`
	AnoValidade string `json:"anoValidade"`
	CVV         string `json:"cvv"`
}

type ClienteDto struct {
	Nome   string    `json:"nome"`
	Email  string    `json:"email"`
	Cartao CartaoDto `json:"cartao"`
}

type CriarPagamentoDto struct {
	PedidoId   string     `json:"pedidoId"`
	ValorTotal int64      `json:"valorTotal"`
	Cliente    ClienteDto `json:"cliente"`
}

type NotificarPagamentoDto struct {
	Id            string                    `json:"id"`
	PedidoId      string                    `json:"pedidoId"`
	ValorTotal    int64                     `json:"valorTotal"`
	Status        pagamento.StatusPagamento `json:"status"`
	TransactionId string                    `json:"transactionId"`
	Erro          string                    `json:"erro"`
	EmailCliente  string                    `json:"emailCliente"`
	CriadoEm      time.Time                 `json:"criadoEm,omitempty"`
}
