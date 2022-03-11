package main

import (
	"fmt"

	"github.com/wendy248/module"
)

func main() {
	nama := "Wendy"
	favorit := "kopi"

	module.Intro()
	fmt.Println(module.Salam(nama))

	hargaSebelum, hargaSesudah := module.Recent(favorit)
	fmt.Println("Harga minuman", favorit, "sebesar :", hargaSebelum, "sehingga perlu membayar (PPN) :", hargaSesudah)
}
