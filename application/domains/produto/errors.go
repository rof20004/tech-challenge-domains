package produto

import "errors"

var (
	ErroIdProdutoInvalido      = errors.New("o id do produto é inválido")
	ErroNomeProdutoObrigatorio = errors.New("o nome do produto é obrigatório")
	ErroPrecoProdutoInvalido   = errors.New("o preço do produto deve ser maior que zero")
	ErroTipoProdutoInvalido    = errors.New("o tipo do produto deve ser: Lanche, Acompanhamento, Bebida ou Sobremesa")
)
