package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 900
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Print("Hesaptaki para yetersiz")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paranız hazırlanıyor..")
		fmt.Println("Dikkat, tüm paranız çekiliyor.")
	} else {
		fmt.Println("Paranız hazırlanıyor.")
	}

}
