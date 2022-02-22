package main

import (
	"fmt"
	"github.com/antonmedv/expr"
)

type Env struct {
	Greet string
	Names []string
	Param Params
}

type Field struct {
	Str string
}

type Params struct {
	Name string

}

func (m Env) Tprint(greet string, params []int) Field {
	return Field{
		Str: fmt.Sprintf(greet, params),
	}
}

func (m Env) TradeNum() int {
	fmt.Printf("ni hao hello\n")
	return 4
}

func main() {
	env := Env{}
	condStr := "2 <= TradeNum() and TradeNum() > 5"
	p, flag := expr.Compile(condStr, expr.Env(env))
	fmt.Printf("flag:%v\n", flag)
	out, err := expr.Run(p, env)

	if err != nil {
		panic(err)
	}
	fmt.Print(out)
}
