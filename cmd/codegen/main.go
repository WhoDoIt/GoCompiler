package main

import (
	"fmt"
	"os"
	"strings"
)

func GenerateLang(name string, subclasses [][]string) {

	f, err := os.Create("internal/syntaxtree/" + strings.ToLower(name) + ".go")
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString("package syntaxtree\n")
	f.WriteString("import (\n\"github.com/WhoDoIt/GoCompiler/internal/tokenizer\"\n)\n")
	f.WriteString("type " + name + " interface {}\n")
	for _, v := range subclasses {
		f.WriteString("type " + v[0] + " struct{\n")
		for j := 1; j < len(v); j += 1 {
			f.WriteString(v[j] + "\n")
		}
		f.WriteString("}\n")
	}
	f.WriteString("type " + name + "Visitor[E any] interface{\n")
	for _, v := range subclasses {
		f.WriteString("Visit" + v[0] + "(" + strings.ToLower(name) + " " + v[0] + ") E\n")
	}
	v := strings.ToLower(name)
	f.WriteString("}\n")
	f.WriteString("func Accept" + name + "[E any](visitor " + name + "Visitor[E], " + strings.ToLower(name) + " " + name + ") E {\n")
	f.WriteString("switch val := " + v + ".(type) {\n")
	for _, v := range subclasses {
		f.WriteString("case " + v[0] + ":\nreturn visitor.Visit" + v[0] + "(val)\n")
	}
	f.WriteString("}\n")
	f.WriteString("return *new(E)\n}\n")
}

func main() {
	GenerateLang("Expr", [][]string{
		{"BinaryExpr", "Left Expr", "Operator tokenizer.Token", "Right Expr"},
		{"UnaryExpr", "Operator tokenizer.Token", "Right Expr"},
		{"GroupingExpr", "Inside Expr"},
		{"CallExpr", "Calle Expr", "Paren tokenizer.Token", "Arguments []Expr"},
		{"LiteralExpr", "Value tokenizer.Token"},
	})
	GenerateLang("Stmt", [][]string{
		{"ExpressionStmt", "Expression Expr"},
		{"PrintStmt", "Expression Expr"},
		{"BlockStmt", "Statements []Stmt"},
		{"IfStmt", "Condition Expr", "Block Stmt"},
		{"VarDeclStmt", "Name tokenizer.Token", "Expression Expr"},
		{"ForStmt", "PreStatement Stmt", "Condition Expr", "PostStatement Expr", "Block Stmt"},
	})
}
