// Package ssa предоставляет функции для построения SSA представления
package ssa

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ssa"
)

// Builder отвечает за построение SSA из исходного кода Go
type Builder struct {
	fset *token.FileSet
}

// NewBuilder создаёт новый экземпляр Builder
func NewBuilder() *Builder {
	return &Builder{
		fset: token.NewFileSet(),
	}
}

// TODO: Реализуйте следующие методы в рамках домашнего задания

// ParseAndBuildSSA парсит исходный код Go и создаёт SSA представление
// Возвращает SSA программу и функцию по имени
func (b *Builder) ParseAndBuildSSA(source string, funcName string) (*ssa.Function, error) {
	// 1. Парсинг исходного кода с помощью go/parser
	f, err := parser.ParseFile(b.fset, "", source, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	conf := &types.Config{
		Importer: importer.Default(),
	}

	pkg, err := conf.Check("main", b.fset, []*ast.File{f}, info)
	if err != nil {
		return nil, fmt.Errorf("type check error: %w", err)
	}

	// 3. Создание SSA-программы
	prog := ssa.NewProgram(b.fset, ssa.SanityCheckFunctions)
	ssaPkg := prog.CreatePackage(pkg, []*ast.File{f}, info, true)
	prog.Build()

	// 4. Поиск нужной функции по имени
	fn := ssaPkg.Func(funcName)
	if fn == nil {
		return nil, fmt.Errorf("function %q not found in package", funcName)
	}

	return fn, err
}
