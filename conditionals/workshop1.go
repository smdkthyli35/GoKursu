package conditionals

import "fmt"

func Workshop1() {
	sayi1 := 10
	sayi2 := 40
	sayi3 := 30

	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}

	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Println("En büyük sayı: ", enBuyuk)

}
