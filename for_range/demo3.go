package for_range

import "fmt"

func Demo3() {
	sozluk := map[string]string{"book": "kitap", "pencil": "kalem"}

	for k := range sozluk {
		fmt.Println(k)
	}
}
