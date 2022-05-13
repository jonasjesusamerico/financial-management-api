package enum

type FormaPagamento int

const (
	DINHEIRO FormaPagamento = iota
	DEBITO   FormaPagamento = iota
	CREDITO  FormaPagamento = iota
	PIX      FormaPagamento = iota
)
