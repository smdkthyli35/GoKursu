package arrays

import "fmt"

func Demo2() {
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "İzmir"
	sehirler[2] = "İstanbul"
	sehirler[3] = "Bursa"
	sehirler[4] = "Muğla"
	// fmt.Print(sehirler)

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}
