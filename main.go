package main

import (
	"fmt"

	"github.com/wendy248/module/v2"
)

func main() {
	var nama, favorit, pilihan string
	var jumlah, hargaSebelum2, hargaDiskon2 float32

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
	fmt.Println("")
	module.Member(nama, anggotaMember)

	//struct
	module.TampilJaga()

	//struct method
	jaga1 := module.DaftarJaga{Nama: "Budi"}
	jaga1.ShiftJaga(nama)

	//struck anonymous
	module.DaftarMenu()

order:
	fmt.Printf("Masukkan nama menu yang ingin dipesan : ")
	fmt.Scan(&favorit)

	fmt.Printf("Masukkan jumlah porsi yang diinginkan : ")
	fmt.Scan(&jumlah)

	//function multiple return
	hargaSebelum, hargaDiskon := module.Recent(favorit, jumlah)
	hargaSebelum2 = hargaSebelum2 + hargaSebelum
	hargaDiskon2 = hargaDiskon2 + hargaDiskon

	fmt.Println("\nTotal harga minuman", favorit, "sebesar :", hargaSebelum2, "dan setelah diskon member (khusus member) :", hargaDiskon2)

	fmt.Printf("\nApakah ingin menambah pesanan? (Y / N) : ")
	fmt.Scan(&pilihan)
	fmt.Println("")

	if pilihan == "Y" || pilihan == "y" {
		goto order
	} else {
		//interface struct
		pjk := module.PPN{Harga: float32(hargaDiskon2)}
		module.Menghitung(pjk.Harga)

	}
}
