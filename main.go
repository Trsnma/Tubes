package main

import (
	"fmt"
)

func main() {
	var harga []float64 
	var diskon int
	// var inputHarga float64

	totalPembelian := 0
	
	fmt.Println(">> Selamat datang di GTR MART, silahkan masukkan harga dari setiap barang yang anda beli\n>> Masukkan harga HARUS menggunakan tanda titik (desimal)\n>> Input angka 0.0 untuk menyelesaikan\n")
	n := 1
	i := 1


	for m := 0; m == n; m++ {
		fmt.Printf("Harga Barang ke %d : ", i)
	    fmt.Scanf("%f\n", &harga[m])

		if harga[m] != 1  {
			fmt.Print("EROR")
		} else if harga[m] < 0.0 {
			fmt.Print("EROR")
		} else if harga[m] == 0.0 {
			m = n
		}

		if harga[m] > 0.0 {
			n++
		}
	}
	
	
	l := len(harga)
	
	for i := 0; i < l - 1; i++ {
		for j := 0; j < l - 1 - i; j++ {
			if harga[j] > harga[j+1] {
				harga[j], harga[j+1] = harga[j+1], harga[j]
			}
		}
	}
	
	// --- Perhitungan Diskon ---
	
if totalPembelian >= 100.000 && totalPembelian < 200.000  {
      diskon = 5
  } else if totalPembelian >= 200.000 && totalPembelian < 300.000 {
      diskon = 10
  } else if totalPembelian >= 300.000 {
      diskon = 20
  }
	
	hasilDiskon := diskon/100
	// Perhitungan sudah natural karena totalPembelian adalah float64
	setelahDiskon := totalPembelian - (totalPembelian*hasilDiskon)
	

	fmt.Println("\n=========================================================================================")
	if harga == nil {
		fmt.Println("\n>> KAMU TIDAK MEMBELI BARANG APAPUN ")
	} else {
		fmt.Println("\nUrutan Harga :", harga)
	}
	

	// 🚨 PERUBAHAN 6: Mencetak Total Pembelian sebagai float64 (%.2f)
	fmt.Printf("Total Pembelian Anda : %.3f\n", totalPembelian)
	fmt.Printf("Diskon yang anda peroleh ( dalam bentuk persen ) : %.0f \n", diskon)
	fmt.Printf("Total pembelian setelah diskon : %.3f\n", setelahDiskon)
	
	
	fmt.Println("\n=========================================================================================")
}