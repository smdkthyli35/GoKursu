package main

import (
	"fmt"
	"golesson/functions"
)

func main() {
	// variables.Demo1()
	// fmt.Print()
	// conditionals.Demo1()
	// conditionals.Workshop1()
	// loops.Demo1()
	// arrays.Demo1()
	//arrays.Demo2()
	// slices.Demo1()
	// slices.Demo2()
	// functions.SelamVer()
	// var sonuc = functions.Topla(2, 3)
	// fmt.Println(sonuc)
	// functions.SelamVer("Samed")

	// sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(20, 5)
	// fmt.Println(sonuc1, sonuc2, sonuc3, sonuc4)

	var toplam = functions.ToplaVariadic(1, 5, 6, 8, 9, 5)
	fmt.Println("Toplam: ", toplam)

	sayilar := []int{10, 15, 65}
	fmt.Println(functions.ToplaVariadic(sayilar...))
}
