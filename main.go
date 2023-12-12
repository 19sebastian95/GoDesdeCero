package main

import (
	"fmt"

	"github.com/19sebastian95/GoDesdeCero/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(132456)
	fmt.Println(estado)
	fmt.Println(texto)
}
