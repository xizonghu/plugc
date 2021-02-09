package pretreat

import (
	"bytes"
	log "github.com/sirupsen/logrus"
)

const (
	fileSizeMax        = 4 * 1024 * 1024
	noteModeNull       = 0
	noteModeSingleLine = 1
	noteModeMultiLine  = 2
)

type PlugcPreparer struct {
}

func (p *PlugcPreparer) removeNotes(text *bytes.Buffer) (*bytes.Buffer, error) {
	var b = text.Bytes()
	n := len(b)
	mode := noteModeNull
	start := 0
	for i := 0; i < n; i++ {
		ch := b[i]

		switch mode {
		case noteModeNull:
			if ch == '/' {
				ch = b[i+1]
				if ch == '/' {
					mode = noteModeSingleLine
					start = i
				} else if ch == '*' {
					mode = noteModeMultiLine
					start = i
				}
			}
			break
		case noteModeSingleLine:
			if ch == '\n' {
				i++
				b = append(b[:start], b[i:]...)
				n = len(b)
				i = start - 1
				mode = noteModeNull
			}
			break
		case noteModeMultiLine:
			if ch == '*' {
				ch = b[i+1]
				if ch == '/' {
					i++
					i++
					b = append(b[:start], b[i:]...)
					n = len(b)
					i = start - 1
					mode = noteModeNull
				}
			}
			break
		}
	}
	return bytes.NewBuffer(b), nil
}
func (p *PlugcPreparer) replaceMacro(text *bytes.Buffer) (*bytes.Buffer, error) {
	return text, nil
}

func (p *PlugcPreparer) addBlank() {}

func (p *PlugcPreparer) Run(text *bytes.Buffer) (out *bytes.Buffer, err error) {

	defer func() {
		if err != nil {
			log.Error(err)
		}
	}()

	if out, err = p.removeNotes(text); err != nil {
		return
	}

	if out, err = p.replaceMacro(out); err != nil {
		return
	}

	return  out, nil
}

func (p *PlugcPreparer) Print(out *bytes.Buffer) string {
	return out.String()
}

func PlugcPreparerNew() *PlugcPreparer {
	return &PlugcPreparer{
	}
}
