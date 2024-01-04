package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestrosSlices() {
	fmt.Println(tablaS)

	procion := arreglo[3:]   // Slice creado desde un vector, desde la posoción 3
	procion2 := arreglo[:5]  //Slice creado desde un vector, desde la posición 0 gasta la 5
	procion3 := arreglo[6:7] //Slice creado desde un vector, desde la posición 6 hasta la 7

	fmt.Println(procion)
	fmt.Println(procion2)
	fmt.Println(procion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Println("Largo", len(elementos), ", Capacidad", cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Println("Largo", len(nums), ", Capacidad", cap(nums))
}
