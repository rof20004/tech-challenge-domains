package dtos

type CartaoDto struct {
	Bandeira    string `json:"bandeira"`
	Numero      string `json:"numero"`
	MesValidade string `json:"mesValidade"`
	AnoValidade string `json:"anoValidade"`
	CVV         string `json:"cvv"`
}

type CriarPagamentoDto struct {
	PedidoId   string `json:"pedidoId"`
	ValorTotal int64  `json:"valorTotal"`
	Cliente    struct {
		Nome   string    `json:"nome"`
		Email  string    `json:"email"`
		Cartao CartaoDto `json:"cartao"`
	} `json:"cliente"`
}