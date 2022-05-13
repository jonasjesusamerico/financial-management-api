package models

import "gorm.io/gorm"

type Estabelecimento struct {
	gorm.Model
	Descricao string
	CNPJ      string
	SetorID   uint
	Setor     Setor
}
