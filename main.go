package main

import (
	"fmt"

	"github.com/wendy248/module"
)

func main() {
	nama := "Wendy"
	favorit := "kopi"
	member := "y"

	//function
	module.Intro()

	//function return
	fmt.Println(module.Salam(nama))
	
	//function multiple return
	hargaSebelum, hargaSesudah := module.Recent(favorit)
	fmt.Println("Harga minuman", favorit, "sebesar :", hargaSebelum, "sehingga perlu membayar (PPN) :", hargaSesudah)

	//anony func
	anggotaMember := func(member string) bool{
		if member == "y" || member == "Y"{
			return true
		} else {
			return false
		}
	}
	module.Member(nama, anggotaMember)
}
