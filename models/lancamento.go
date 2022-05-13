package models

import (
	"gorm.io/gorm"
)

type Lancamento struct {
	gorm.Model
	Descricao string `json:"descricao"`
	ContaID   uint   `json:"-"`
	Conta     Conta  `json:"conta"`
	// Setor           Setor
	// FormaPagamento  enum.FormaPagamento
	// Valor           float64
	// DataCompra      time.Time
	// Estabelecimento Estabelecimento
}

// func (l Lancamento) SituacaoCompra() enum.TipoConta {
// 	if l.Conta.TipoConta == enum.TIPO_CONTA_CORRENTE {
// 		return enum.TIPO_CONTA_CORRENTE
// 	}
// 	return enum.TIPO_CONTA_DEBITO
// }
