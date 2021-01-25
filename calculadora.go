package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct {
}

func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)

	operador1 := parse(entradaLimpia[0])
	operador2 := parse(entradaLimpia[1])

	switch operador {
	case "+":
		return operador1 + operador2
	case "-":
		return operador1 - operador2
	case "*":
		return operador1 * operador2
	case "/":
		return operador1 / operador2
	default:
		fmt.Println("El operador " + operador + " no está soportado")
		return 0
	}
}

func parse(entrada string) int {
	operador, _ := strconv.Atoi(entrada)

	return operador
}

func ReadEntry() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	operacion := scanner.Text()

	return operacion
}

/*func main() {
	entrada := readEntry()
	operador := readEntry()
	c := calc{}

	fmt.Println(entrada)
	fmt.Println(operador)
	fmt.Println(c.operate(entrada, operador))

	// calc.operate(entrada, operador)
}*/
