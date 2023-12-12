package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo32 float32
var Sueldo64 float64
var Fecha time.Time

func RestoVariables() {
	Nombre = "Sebastian"
	Estado = true
	Sueldo32 = 1577.6
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo32)
	fmt.Println(Fecha)
}

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
