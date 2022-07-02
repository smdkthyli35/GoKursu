package variables

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya!"

	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)

	var metin1 int = 15
	fmt.Println(metin1)

	var metin2 float32 = 0.5
	fmt.Println(metin2)

	metin3 := 25
	fmt.Println(metin3)
	fmt.Printf("veri tipi : %T", metin3)
}
