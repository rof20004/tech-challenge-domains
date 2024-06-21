package pagamento

import (
	"github.com/rof20004/tech-challenge-domains/application/utils"
	"time"

	"github.com/google/uuid"
)

type StatusPagamento string

type Pagamento struct {
	Id            string          `json:"id"`
	TransactionId string          `json:"transactionId"`
	PedidoId      string          `json:"pedidoId"`
	Valor         int64           `json:"valor"`
	EmailCliente  string          `json:"emailCliente"`
	Status        StatusPagamento `json:"status"`
	MotivoErro    string          `json:"motivoErro"`
	CriadoEm      time.Time       `json:"criadoEm"`
	AtualizadoEm  time.Time       `json:"atualizadoEm"`
}

func NovoPagamento(pedidoId, emailCliente string, valor int64) (Pagamento, error) {
	p := Pagamento{
		Id:           uuid.NewString(),
		PedidoId:     pedidoId,
		Valor:        valor,
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

	if p.Valor < 1 {
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

func (p *Pagamento) Atualizar(status StatusPagamento) error {
	p.Status = status
	p.AtualizadoEm = time.Now().UTC()
	return p.validar()
}
