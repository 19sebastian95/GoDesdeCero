package iteraciones

import "fmt"

func IterarFor() {

	for i := 0; i <= 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
