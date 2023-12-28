package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplicacion() string {
	var numero int
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	var texto string

	for {
		fmt.Println("Ingrese número: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintln(numero, "x", i, "=", numero*i)
	}

	// for i := 1; i <= 10; i++ {
	// 	texto2 += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
	// }

	return texto
}
