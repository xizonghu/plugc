package compile

const (
	ERROR_NULL		= iota
)

const (
	ERROR_FILE_SIZE					= "file size error"
)

type PlugcError struct {
	code int
	errMsg string
}

func PlugErrorNew(code int) {

}