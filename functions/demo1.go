package functions

import "fmt"

func Topla(sayi1 int, sayi2 int) int {
	var toplam int = sayi1 + sayi2
	return toplam
}

func SelamVer(kullanici string) {
	fmt.Println("Merhaba", kullanici)
}
