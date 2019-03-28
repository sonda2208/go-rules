package rules

type Op string

const (
	AND Op = "&&"
	OR  Op = "||"
	EQ  Op = "=="
	NEQ Op = "!="
	LT  Op = "<"
	LTE Op = "<="
	GT  Op = ">"
	GTE Op = ">="
)
