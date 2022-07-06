package structs

import "fmt"

func Demo1() {
	fmt.Println(product{name: "Bilgisayar", unitPrice: 500, brand: "Lenovo"})
}

type product struct {
	name         string
	unitPrice    float64
	brand        string
	discountRate int
}
