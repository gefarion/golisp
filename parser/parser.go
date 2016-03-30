package main

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	"golisp/types"
)

func addAtom(sexpr *string, stack_expr []types.LispList) {
	if len(*sexpr) > 0 {
		stack_expr[len(stack_expr)-1] = append(stack_expr[len(stack_expr)-1], parseAtom(*sexpr))
		*sexpr = ""
	}
}

func parseAtom(sexpr string) types.LispExpression {
	num, err := strconv.Atoi(sexpr)
	if err == nil {
		return types.LispInt(num)
	} else {
		return types.LispString(sexpr)
	}
}

func LispExprToSExpr(lexpr types.LispExpression) string {
	return strings.Replace(
		strings.Replace(fmt.Sprintf("%v", lexpr),
		"[", "(", -1),"]", ")", -1)
}

func SExprToLispExpr(sexpr string) (types.LispExpression, error) {

	stack_expr, current_sexpr := []types.LispList{ types.LispList{} }, ""
	
	for _, char := range strings.Split(strings.TrimSpace(sexpr) + " ", "") {
		switch char {
			case "(":
				stack_expr = append(stack_expr, types.LispList{})
			case ")":
				addAtom(&current_sexpr, stack_expr)
				
				last := stack_expr[len(stack_expr)-1]
				stack_expr = stack_expr[:len(stack_expr)-1]
				
				if len(stack_expr) == 0 {
					return nil, errors.New("Unbalanced parentesis")
				}
				stack_expr[len(stack_expr)-1] = append(stack_expr[len(stack_expr)-1], last)
			case " ":
				addAtom(&current_sexpr, stack_expr)
			default:
				current_sexpr += char
		}
	}
	
	if len(stack_expr) == 1 {
		return stack_expr[0][0], nil
	} else {
		return nil, errors.New("Unbalanced parentesis")
	}
	
}

func main() {
	// Pasar a tests de unidad
	exp, err := SExprToLispExpr("(1 sda)    32 asd 2)")
	fmt.Println(LispExprToSExpr(exp))
	fmt.Println(exp)
	fmt.Println(err)
}
