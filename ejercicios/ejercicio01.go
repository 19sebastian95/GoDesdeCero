package ejercicios

import (
	"fmt"
	"strconv"
)

func DevolverValores(valor string) (int, string) {

	valorEntero, err := strconv.Atoi(valor)

	if valorEntero > 100 {
		fmt.Println("Es mayor a 100", valorEntero)
	} else {
		fmt.Println("Es menor a 100")
	}

	return valorEntero, err.Error()
}
