package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Digite a operação (ex: 2 + 2): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.ReplaceAll(strings.TrimSpace(input), " ", "")

	var operator string
	for _, op := range []string{"+", "-", "*", "/"} {
		if strings.Contains(input, op) {
			operator = op
			break
		}
	}

	if operator == "" {
		panic("Operação inválida")
	}

	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		panic("Formato inválido")
	}

	num1, err1 := strconv.Atoi(parts[0])
	num2, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		panic("Número inválido")
	}

	result := getResult(num1, num2, operator)

	fmt.Printf("%d %s %d = %d\n", num1, operator, num2, result)
}

func getResult(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("Divisão por zero")
		}
		return num1 / num2
	default:
		panic("Operação inválida")
	}
}
