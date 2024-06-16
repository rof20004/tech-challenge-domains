package produto

import (
	"strings"
	"time"

	"github.com/rof20004/tech-challenge-domains/application/utils"
)

type TipoProduto string

type Produto struct {
	Id           string      `json:"id"`
	Nome         string      `json:"nome"`
	Descricao    string      `json:"descricao"`
	Preco        int64       `json:"preco"`
	Tipo         TipoProduto `json:"tipo"`
	CriadoEm     time.Time   `json:"criadoEm,omitempty"`
	AtualizadoEm time.Time   `json:"atualizadoEm,omitempty"`
}

func NovoProduto(nome, descricao string, preco int64, tipo TipoProduto) (Produto, error) {
	p := Produto{
		Id:        utils.GenerateUuid(),
		Nome:      nome,
		Descricao: descricao,
		Preco:     preco,
		Tipo:      tipo,
		CriadoEm:  time.Now().UTC(),
	}

	return p, p.validar()
}

func (p *Produto) validar() error {
	if strings.TrimSpace(p.Nome) == "" {
		return ErroNomeProdutoObrigatorio
	}

	if p.Preco < 1 {
		return ErroPrecoProdutoInvalido
	}

	if p.Tipo != LANCHE &&
		p.Tipo != ACOMPANHAMENTO &&
		p.Tipo != BEBIDA &&
		p.Tipo != SOBREMESA {
		return ErroTipoProdutoInvalido
	}

	return nil
}

func (p *Produto) Atualizar(nome, descricao string, preco int64, tipo TipoProduto) error {
	p.Nome = nome
	p.Descricao = descricao
	p.Preco = preco
	p.Tipo = tipo
	p.AtualizadoEm = time.Now().UTC()
	return p.validar()
}
