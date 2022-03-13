package main

import (
	"fmt"

	"github.com/wendy248/module"
)

func main() {
	var nama, favorit string
	var jumlah float32

	//function
	module.Intro()

	fmt.Printf("Masukkan nama anda : ")
	fmt.Scan(&nama)

	//function return
	fmt.Println(module.Salam(nama))

	//anonymous func
	anggotaMember := func(nama string) bool {
		if nama == "Wendy" || nama == "wendy" {
			return true
		} else if nama == "Harry" || nama == "harry" {
			return true
		} else {
			return false
		}
	}
	module.Member(nama, anggotaMember)

	//struct
	module.TampilJaga()

	//struct method
	jaga1 := module.DaftarJaga{Nama: "Budi"}
	jaga1.ShiftJaga(nama)

	//struck anonymous
	module.DaftarMenu()

	fmt.Printf("Masukkan nama menu yang ingin dipesan : ")
	fmt.Scan(&favorit)

	fmt.Printf("Masukkan jumlah porsi yang diinginkan : ")
	fmt.Scan(&jumlah)

	//function multiple return
	hargaSebelum, hargaDiskon := module.Recent(favorit, jumlah)
	fmt.Println("\nHarga minuman", favorit, "sebesar :", hargaSebelum, "dan setelah diskon member (khusus member) :", hargaDiskon)

	//interface struct
	pjk := module.PPN{Harga: float32(hargaSebelum)}
	module.Menghitung(pjk.Harga)

}
