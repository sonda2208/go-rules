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
	IN  Op = "in"
	ADD Op = "+"
	SUB Op = "-"
	MUL Op = "*"
	DIV Op = "/"
	MOD Op = "%"
)
