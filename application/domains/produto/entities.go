package produto

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"strings"
	"time"

	"github.com/rof20004/tech-challenge-domains/application/utils"
)

type TipoProduto string

type Produto struct {
	Id           string      `json:"id" bson:"id"`
	Nome         string      `json:"nome" bson:"nome"`
	Descricao    string      `json:"descricao" bson:"descricao"`
	Preco        int64       `json:"preco" bson:"preco"`
	Tipo         TipoProduto `json:"tipo" bson:"tipo"`
	CriadoEm     time.Time   `json:"criadoEm,omitempty" bson:"criadoEm"`
	AtualizadoEm time.Time   `json:"atualizadoEm,omitempty" bson:"atualizadoEm"`
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

	return p, p.Validar()
}

func (p *Produto) Validar() error {
	if _, err := uuid.Parse(p.Id); err != nil {
		return errors.Wrap(ErroIdProdutoInvalido, err.Error())
	}

	if strings.TrimSpace(p.Nome) == "" {
		return ErroNomeProdutoObrigatorio
	}

	if p.Preco < 1 {
		return ErroPrecoProdutoInvalido
	}

	if p.Tipo != Lanche &&
		p.Tipo != Acompanhamento &&
		p.Tipo != Bebida &&
		p.Tipo != Sobremesa {
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
	return p.Validar()
}
