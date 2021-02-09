package stmt

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"plugc/compiler/ast/node"
	"plugc/compiler/errors"
	"plugc/compiler/token"
	"plugc/compiler/token/utils"
	"reflect"
)

type SegmentStmt struct {
	Stmts []IStmt
}

func (p *SegmentStmt) Check(reader token.ITokenReader) (ok bool) {
	if tk := reader.Read(); utils.IsBraceLeft(tk) {
		ok = true
	}
	return
}

func (p *SegmentStmt) Parse(reader token.ITokenReader) (root node.IAstNode, err error) {
	log.Debugf("parse segment stmt at: line %d", reader.Read().LineNumber)
	if tk := reader.Read(); !utils.IsBraceLeft(tk) {
		err = errors.New(fmt.Sprintf("expect '{': %s", tk))
		return
	}

	reader.Next()

	seg := node.SegmentNew()
	root = seg

	for tk := reader.Read(); !utils.IsBraceRight(tk); tk = reader.Read() {
		var statement IStmt
		for _, item := range Stmts {
			if item.Check(reader) {
				statement = item
				break
			}
		}

		if !reflect.ValueOf(statement).IsValid() {
			err = errors.New(fmt.Sprintf("can not parse: %s", reader))
			return
		}

		var nod node.IAstNode
		if nod, err = statement.Parse(reader); err != nil {
			return
		}
		seg.Stmts = append(seg.Stmts, nod)
	}

	if tk := reader.Read(); !utils.IsBraceRight(tk) {
		err = errors.New(fmt.Sprintf("expect '}': %s", tk))
	}
	reader.Next()

	return
}

func SegmentStmtNew() *SegmentStmt {
	return &SegmentStmt{}
}
