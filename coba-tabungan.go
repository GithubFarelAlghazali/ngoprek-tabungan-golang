package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func input(prompt string)string{
	fmt.Print(prompt)
	teks,_ := reader.ReadString('\n')
	return strings.TrimSpace(teks) 
}

type tabungan struct{
	nama string
	target float64
	terkumpul float64
	status string
}


func tampilkanMenu(){
	fmt.Println("Daftar perintah: ");
	fmt.Println("1 Tampilkan daftar tabungan");
	fmt.Println("2 Buat tabungan");
	fmt.Println("3 Edit tabungan");
	fmt.Println("4 Hapus tabungan");
	fmt.Println("5 Isi tabungan");
	fmt.Println("6 Ambil isi tabungan");
	fmt.Println("0 Keluar");
}

func formatRupiah(num float64) string {
	s := strconv.FormatFloat(num, 'f', 0, 64)
	var bagianBagian []string
	for len(s) > 3{
		bagianBagian = append([]string{s[len(s)-3:]}, bagianBagian...)
		s = s[:len(s)-3]
	}
	bagianBagian = append([]string{s}, bagianBagian...)
	return "Rp " + strings.Join(bagianBagian,".")
}

func tampilkanTabungan(daftar []tabungan){
	if len(daftar) == 0{
		fmt.Println("Belum ada tabungan")
	} else {	
		fmt.Printf("%-4s %-20s %-18s %-18s %-18s %-8s\n", "No", "Nama", "Target", "Terkumpul", "Kekurangan", "Status")
		fmt.Println(strings.Repeat("=", 100))
		for i, v := range daftar {
			var kekurangan string
			if(v.terkumpul < v.target){
				kekurangan = formatRupiah(v.target - v.terkumpul)
			} else {
				kekurangan = "penuh"
			}
			
			fmt.Printf("%-4d %-20s %-18s %-18s %-18s %-8s\n", i+1, v.nama, formatRupiah(v.target),formatRupiah(v.terkumpul), kekurangan, v.status)
		}
		fmt.Println()
	}
	
}

func buatTabunganBaru(daftar []tabungan,namaTabungan string, targetTabungan float64)[]tabungan{
	baru := tabungan{nama: namaTabungan, target: targetTabungan, terkumpul: 0, status: "progress"}
	fmt.Printf("\nTabungan %q berhasil dibuat!\n", namaTabungan)
	fmt.Println()
	return append(daftar, baru)
}

func hapusTabungan(namaTabungan string, daftar []tabungan)[]tabungan{
	for i,v:= range daftar{
		if(strings.EqualFold(v.nama,namaTabungan)){
			fmt.Printf("Tabungan %q dihapus\n", v.nama)
			return append(daftar[:i], daftar[i+1:]...)
		}
	}
	fmt.Print("Tabungan tidak ditemukan")
	return daftar
}

func editTabungan(daftar []tabungan,namaTabungan string) {
	for i, v:= range daftar{
		if(strings.EqualFold(v.nama, namaTabungan)){
			namaBaru := input("Masukkan nama baru (enter untuk melewati) : ")
			targetBaru, err := strconv.ParseFloat( input("Masukkan target baru : Rp "), 64)
			if err != nil {
				fmt.Println("Target harus berupa angka!")
				return
			}
			if(namaBaru == ""){
				daftar[i].nama = namaBaru
			}
			
			daftar[i].target = targetBaru
			fmt.Println("Tabungan berhasil diupdate!")
			return
		}
	}
	fmt.Println("Tabungan tidak ditemukan!")
}

func isiTabungan(daftar []tabungan, namaTabungan string) {
	for i, v:= range daftar{
		if(strings.EqualFold(v.nama, namaTabungan)){
			isiBaru, err := strconv.ParseFloat( input("Masukkan nominal pengisian : Rp "), 64)
			if err != nil {
				fmt.Println("Target harus berupa angka!")
				return
			}
			daftar[i].terkumpul = isiBaru + v.terkumpul
			if(daftar[i].terkumpul >= daftar[i].target){
				daftar[i].status = "tercapai"
			}
			fmt.Println("Tabungan berhasil diupdate!")
			return
		}
	}
	fmt.Println("Tabungan tidak ditemukan!")
}

func ambilTabungan(daftar []tabungan, namaTabungan string) {
	for i, v:= range daftar{
		if(strings.EqualFold(v.nama, namaTabungan)){
			nominal, err := strconv.ParseFloat( input("Masukkan nominal pengambilan : Rp "), 64)
			fmt.Println()
			if err != nil {
				fmt.Println("Target harus berupa angka!")
				return
			}
			if(nominal > v.terkumpul){
				fmt.Println("Saldo tabungan tidak cukup untuk diambil")
				return
			}
			daftar[i].terkumpul = daftar[i].terkumpul - nominal
			fmt.Println("Tabungan berhasil diupdate!")
			fmt.Println()
			return
		}
	}
	fmt.Println()
	fmt.Println("Tabungan tidak ditemukan!")
	fmt.Println()
}


func main(){
	daftarTabungan := []tabungan{}

	fmt.Println("Aplikasi tabungan")
	for{
		tampilkanMenu()
		
		pilihan, _ := strconv.Atoi(input("Masukkan nomor operasi : "))
		
		if pilihan == 0 {
			break
		}
		
		switch pilihan{
		case 1:
			tampilkanTabungan(daftarTabungan)
		case 2:
			nama:= input("Masukkan nama tabungan baru : ")			
			target, err := strconv.ParseFloat(input("Masukkan jumlah target : Rp "), 64)
			if err == nil{
				daftarTabungan= buatTabunganBaru(daftarTabungan, nama, target)
			}
		case 3:
			nama := input("Masukkan nama tabungan yang ingin diedit : ")
			editTabungan(daftarTabungan, nama)
		case 4:
			nama:= input("Masukkan nama tabungan yang ingin dihapus : ")
			daftarTabungan = hapusTabungan(nama, daftarTabungan)
		case 5:
			nama:= input("Masukkan nama tabungan yang ingin diisi : ")
			isiTabungan(daftarTabungan, nama)
		case 6:
			nama:= input("Masukkan nama tabungan yang ingin diambil : ")
			ambilTabungan(daftarTabungan, nama)
		}

	}
}