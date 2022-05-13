package models

import (
	"gorm.io/gorm"
)

type Conta struct {
	gorm.Model
	// TipoConta enum.TipoConta
	NomeConta string
}
