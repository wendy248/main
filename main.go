package main

import (
	"fmt"

	"github.com/wendy248/module"
)

func main() {
	nama := "Wendy"
	favorit := "kopi"

	//function
	module.Intro()

	//function return
	fmt.Println(module.Salam(nama))

	//function multiple return
	hargaSebelum, hargaSesudah := module.Recent(favorit)
	fmt.Println("Harga minuman", favorit, "sebesar :", hargaSebelum, "sehingga perlu membayar (PPN) :", hargaSesudah)

	//anonymous func
	anggotaMember := func(nama string) bool {
		if nama == "Wendy" {
			return true
		} else if nama == "Harry" {
			return true
		} else {
			return false
		}
	}
	module.Member(nama, anggotaMember)
	// module.TampilJaga()

	jaga1 := module.DaftarJaga{module.Nama: "Rian"}
	jaga1.module.Shift(nama)
}
