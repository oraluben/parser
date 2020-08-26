// +build gofuzz

package fuzz

import (
	"strings"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"

	_ "github.com/pingcap/parser/test_driver"
)

func Fuzz(input []byte) int {
	var b strings.Builder

	b.Write(input)

	stmts, err := parse(b.String())
	if err == nil && len(stmts) > 0 {
		return 1
	}
	return 0
}

func parse(sql string) ([]ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return stmtNodes, nil
}
