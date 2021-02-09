package token


type RegexResult struct {
	Str string
	Start uint32
	Stop uint32
}

func RegexCompile(reader TokenReader , expr string) (results []*RegexResult, err error) {
	return
}