package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman Sayısı:", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)
	delete(sozluk, "book")
	fmt.Println("Eleman Sayısı:", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)

	deger, varMi := sozluk["tabl"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu: ", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)
}
