package pagamento

import (
	"strings"
	"time"

	"github.com/rof20004/tech-challenge-domains/application/utils"

	"github.com/google/uuid"
)

type StatusPagamento string

type Pagamento struct {
	Id            string          `json:"id" bson:"id"`
	TransactionId string          `json:"transactionId" bson:"transactionId"`
	PedidoId      string          `json:"pedidoId" bson:"pedidoId"`
	ValorTotal    int64           `json:"valorTotal" bson:"valorTotal"`
	EmailCliente  string          `json:"emailCliente" bson:"emailCliente"`
	Status        StatusPagamento `json:"status" bson:"status"`
	MotivoErro    string          `json:"motivoErro" bson:"motivoErro"`
	CriadoEm      time.Time       `json:"criadoEm" bson:"criadoEm"`
	AtualizadoEm  time.Time       `json:"atualizadoEm" bson:"atualizadoEm"`
}

func NovoPagamento(pedidoId, emailCliente string, valorTotal int64) (Pagamento, error) {
	p := Pagamento{
		Id:           uuid.NewString(),
		PedidoId:     pedidoId,
		ValorTotal:   valorTotal,
		EmailCliente: emailCliente,
		Status:       Pendente,
		CriadoEm:     time.Now().UTC(),
	}

	return p, p.validar()
}

func (p *Pagamento) validar() error {
	if _, err := uuid.Parse(p.PedidoId); err != nil {
		return ErroPedidoIdPagamentoInvalido
	}

	if p.ValorTotal < 1 {
		return ErroValorPagamentoInvalido
	}

	if p.Status != Concluido &&
		p.Status != Erro &&
		p.Status != Pendente {
		return ErroStatusPagamentoInvalido
	}

	if !utils.IsEmailValido(p.EmailCliente) {
		return ErroEmailClientePagamentoInvalido
	}

	return nil
}

func (p *Pagamento) Atualizar(transactionId string, erro error, status StatusPagamento) error {
	if erro == nil && strings.TrimSpace(transactionId) == "" {
		return ErroTransactionIdObrigatorio
	}

	if erro != nil {
		p.MotivoErro = erro.Error()
	}

	p.TransactionId = transactionId
	p.Status = status
	p.AtualizadoEm = time.Now().UTC()

	return p.validar()
}
