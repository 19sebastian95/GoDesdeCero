package main

import (
	"fmt"

	"github.com/19sebastian95/GoDesdeCero/ejercicios"
)

func main() {
	// estado, texto := variables.ConviertoATexto(132456)
	// fmt.Println(estado)
	// fmt.Println(texto)

	// if os := runtime.GOOS; os == "Linux." || os == "OS X." {
	// 	fmt.Println("Esto no es Windows")
	// } else {
	// 	fmt.Println("Esto es Windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Linux")
	// default:
	// 	fmt.Printf("%s \n", os)
	//}

	entero, err := ejercicios.DevolverValores("100")
	fmt.Println(entero)
	if err != "" {
		fmt.Println(err)
	}
}
