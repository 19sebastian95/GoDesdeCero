package main

import (
	"github.com/19sebastian95/GoDesdeCero/middleware"
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

	// entero, texto := ejercicios.DevolverValores("500")
	// fmt.Println(entero)
	// fmt.Println(texto)

	//teclado.IngresoNumeros()
	//iteraciones.IterarFor()

	// ejercicios.Multiplicacion()

	//files.SumaTabla()

	//files.LeoArchivo()

	//funciones.Exponencia(2)

	// mapas.MostrarMapas()

	//users.AltaUsuario()

	// Pedro := new(modelos.Hombre)
	// ejer_interfaces.HumanosRespirando(Pedro)

	//deferpanic.EjemploPanic()

	// canal1 := make(chan bool)
	// go goroutines.MiNombreLento("Sebastian", canal1)
	// fmt.Println("Estoy aqui")
	// defer func() { <-canal1 }()
	// var x string
	// fmt.Scanln(&x)

	//webserver.MiWebServer()

	middleware.MiMiddleware()
}
