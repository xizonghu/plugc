package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"plugc/compiler/ast/node"
	"plugc/compiler/coder"
	"plugc/compiler/lexer"
	"plugc/compiler/parser"
	"plugc/compiler/pretreat"
	"plugc/compiler/token"
	"plugc/module/logutils"
)

func init() {
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&logutils.LogFormatter{})
}

func main() {
	dir := "compiler/examples/test/src/"
	filename := "example00.c"
	fin, _ := os.Open(dir + filename)
	defer fin.Close()

	b, _ := ioutil.ReadAll(fin)

	var err error
	defer func() {
		if err != nil {
			fmt.Printf("err=%s\n", err)
		}
	}()

	var in, out *bytes.Buffer
	var fout *os.File

	in = bytes.NewBuffer(b)

	pre := pretreat.PlugcPreparerNew()
	if out, err = pre.Run(in); err != nil {
		return
	}
	fout, _ = os.Create("build/" + filename + ".pretreat.txt")
	fout.Write([]byte(pre.Print(out)))
	fout.Close()

	var tokens []*token.Token
	lex := lexer.LexerNew()
	if tokens, err = lex.Run(out); err != nil {
		return
	}
	fout, _ = os.Create("build/"+ filename + ".lexer.txt")
	fout.Write([]byte(lex.Print(tokens)))
	fout.Close()

	var root node.IAstNode
	par := parser.CParserNew()
	if root, err = par.Run(tokens); err != nil {
		return
	}
	fout, _ = os.Create("build/"+ filename + ".ast.txt")
	fout.Write([]byte(par.Print(root)))
	fout.Close()

	var bytes []byte
	cod := coder.CoderNew()
	if bytes, err = cod.Run(root); err != nil {
		return
	}

	fout, _ = os.Create("build/"+ filename + ".plc")
	fout.Write(bytes)
	fout.Close()

	fmt.Printf("done\n")
}
