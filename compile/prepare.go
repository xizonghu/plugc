package compile

import (
	"bytes"
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	fileSizeMax        = 4 * 1024 * 1024
	noteModeNull       = 0
	noteModeSingleLine = 1
	noteModeMultiLine  = 2
)

type PlugcPreparer struct {
	inputPath  string
	inputFile  *os.File
	outputFile *os.File
	text       string
	b          []byte
}

func (p *PlugcPreparer) removeNotes() error {
	var b = p.b
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
	p.b = b
	return nil
}
func (p *PlugcPreparer) replaceMacro() error {
	return nil
}

func (p *PlugcPreparer) removeSpaces() error {
	p.b = bytes.ReplaceAll(p.b, []byte{' '}, nil)
	p.b = bytes.ReplaceAll(p.b, []byte{'\r'}, nil)
	p.b = bytes.ReplaceAll(p.b, []byte{'\n'}, nil)
	return nil
}

func (p *PlugcPreparer) addBlank() {}

func (p *PlugcPreparer) Run() error {
	var fp *os.File
	var err error

	defer func() {
		if err != nil {
			log.Error(err)
		}
	}()

	fp, err = os.Open(p.inputPath)
	if err != nil {
		return err
	}
	finfo, _ := os.Stat(p.inputPath)
	size := finfo.Size()
	p.b = make([]byte, size)
	var n int
	if n, err = fp.Read(p.b); err != nil {
		return err
	}
	if n != int(size) {
		return errors.New(ERROR_FILE_SIZE)
	}

	if err = p.removeNotes(); err != nil {
		return err
	}
	if err = p.replaceMacro(); err != nil {
		return err
	}
	if err = p.removeSpaces(); err != nil {
		return err
	}
	if err = p.Save(); err != nil {
		return err
	}

	return nil
}

func (p *PlugcPreparer) Save() error {
	fp, err := os.Create(p.inputPath + ".1")
	if err != nil {
		return err
	}
	fp.Write(p.b)
	fp.Close()
	return nil
}

func PlugcPreparerNew(inputPath string) *PlugcPreparer {
	return &PlugcPreparer{
		inputPath:  inputPath,
		inputFile:  nil,
		outputFile: nil,
		text:       "",
	}
}
