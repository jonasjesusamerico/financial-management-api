package models

import "gorm.io/gorm"

type Carteira struct {
	gorm.Model
	Descricao string
	Valor     float64
}
