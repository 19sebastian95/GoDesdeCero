package deferpanic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es el primer mensaje")
	defer fmt.Println("Este es el mensaje final") // lo ultimo que se va a ejecutar por el defer
	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrió un error que generó Panic \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
