package main

import "fmt"

func main() {
	nama := "Wendy"
	favorit := "kopi"

	Intro()
	fmt.Println(Salam(nama))
	
	hargaSebelum, hargaSesudah := Recent(favorit)
	fmt.Println("Harga minuman :", hargaSebelum, "sehingga perlu membayar (PPN) :", hargaSesudah)
}
