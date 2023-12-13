package ejercicios

import (
	"strconv"
)

func DevolverValores(valor string) (int, string) {

	valorEntero, err := strconv.Atoi(valor)

	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}

	if valorEntero > 100 {
		return valorEntero, "Es mayor a 100"
	} else {
		return valorEntero, "Es menor a 100"
	}
}
