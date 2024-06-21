package pedido

import (
	"time"

	"github.com/rof20004/tech-challenge-domains/application/domains/cliente"
	"github.com/rof20004/tech-challenge-domains/application/domains/produto"
	"github.com/rof20004/tech-challenge-domains/application/utils"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type StatusPedido string

type Pedido struct {
	Id           string            `json:"id"`
	Cliente      cliente.Cliente   `json:"cliente"`
	Produtos     []produto.Produto `json:"produtos"`
	Status       StatusPedido      `json:"status"`
	ValorTotal   int64             `json:"valorTotal"`
	CriadoEm     time.Time         `json:"criadoEm"`
	AtualizadoEm time.Time         `json:"atualizadoEm"`
}

func NovoPedido(cliente cliente.Cliente, produtos []produto.Produto) (Pedido, error) {
	p := Pedido{
		Id:         utils.GenerateUuid(),
		Cliente:    cliente,
		Produtos:   produtos,
		Status:     Recebido,
		ValorTotal: getValorTotal(produtos),
		CriadoEm:   time.Now().UTC(),
	}

	return p, p.validar()
}

func (p *Pedido) validar() error {
	if _, err := uuid.Parse(p.Id); err != nil {
		return errors.Wrap(ErroIdPedidoInvalido, err.Error())
	}

	if err := p.Cliente.Validar(); err != nil {
		return errors.Wrap(ErroClientePedidoInvalido, err.Error())
	}

	if err := p.validarProdutos(); err != nil {
		return errors.Wrap(ErroProdutosPedidoInvalidos, err.Error())
	}

	if p.Status != Recebido &&
		p.Status != EmPreparacao &&
		p.Status != Pronto &&
		p.Status != Finalizado &&
		p.Status != Cancelado {
		return ErroStatusPedidoInvalido
	}

	if p.ValorTotal < 1 {
		return ErroValorTotalPedidoInvalido
	}

	return nil
}

func getValorTotal(produtos []produto.Produto) int64 {
	var total int64

	for _, p := range produtos {
		total += p.Preco
	}

	return total
}

func (p *Pedido) Atualizar(status StatusPedido) error {
	p.Status = status
	p.AtualizadoEm = time.Now().UTC()
	return p.validar()
}

func (p *Pedido) validarProdutos() error {
	for _, pr := range p.Produtos {
		if err := pr.Validar(); err != nil {
			return err
		}
	}

	return nil
}
