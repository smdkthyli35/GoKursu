package slices

import "fmt"

func Demo1() {

	isimler := make([]string, 3)

	fmt.Println(isimler)
	isimler[0] = "Samed"
	isimler[1] = "Samed"
	isimler[2] = "Samed"
	isimler = append(isimler, "Veli")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
