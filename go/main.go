package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jackjohn7/postfix_evaluators/go/stack"
	"github.com/jackjohn7/postfix_evaluators/go/symbols"
)

func main() {
	expression := "3 5 8 + *"

	result := evaluatePostfix(expression)

	fmt.Println(result)
}

func evaluatePostfix(expr string) int {
	t := symbols.New(expr)
	s := stack.New(t.GetSymCount())

	current := t.ReadNext()
	for current != nil {
		var err error
		var val int

		switch *current {
		case "+":
			val, err = arithmetic(s.Pop(), s.Pop(), '+')
			s.Push(val)
		case "-":
			val, err = arithmetic(s.Pop(), s.Pop(), '-')
			s.Push(val)
		case "*":
			val, err = arithmetic(s.Pop(), s.Pop(), '*')
			s.Push(val)
		case "/":
			val, err = arithmetic(s.Pop(), s.Pop(), '/')
			s.Push(val)
		default:
			val, err := strconv.Atoi(*current)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			s.Push(val)
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		current = t.ReadNext()
	}

	result := s.Pop()
	return result
}

func arithmetic(v1 int, v2 int, operator rune) (int, error) {
	switch operator {
	case '+':
		return v1 + v2, nil
	case '-':
		return v1 - v2, nil
	case '*':
		return v1 * v2, nil
	case '/':
		return v1 / v2, nil
	}

	return -1, errors.New("Invalid operator")
}
