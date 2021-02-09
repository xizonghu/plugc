package pretreat

import "bytes"

type IPretreat interface {
	Run(in *bytes.Buffer) (out *bytes.Buffer, err error)
	Print(out *bytes.Buffer) string
}
