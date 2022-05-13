package models

import "gorm.io/gorm"

type Setor struct {
	gorm.Model
	Descricao string
}
