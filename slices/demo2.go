package slices

import "fmt"

func Demo2() {

	sehirler := []string{"İzmir", "İstanbul", "Ankara"}
	fmt.Println(sehirler)
	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)
	sehirler = append(sehirler, "Bursa")
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler[1:3])
}
