package enum

type TipoConta int

const (
	TIPO_CONTA_CORRENTE TipoConta = iota
	TIPO_CONTA_DEBITO   TipoConta = iota
)
