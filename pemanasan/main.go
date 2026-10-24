package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	var nama string = "Aldo"
	var umur int = 20
	kota := "Solo"
	const appName = "cobaCoba"
	paragraf := `Halo gais welcome to my app,
makasi udah dirun`

	fmt.Println("Alo0 reKkSsSS")
	fmt.Printf("Nama aplikasi : %v\n",appName)
	nama = "Dion"
	fmt.Println("Nama : ", nama)
	fmt.Println("Umur : ", umur)
	fmt.Println("Kota : ", kota)
	fmt.Println("Normal : ",paragraf)
	fmt.Println("Kecil : ", strings.ToLower(paragraf))
	fmt.Println("Gede : ", strings.ToUpper(paragraf))
	fmt.Println("Dimulai dengan Hai? : ", strings.HasPrefix(paragraf, "Hai"))
	fmt.Println("Ada kata gais? : ", strings.Contains(paragraf, "gais"))
	fmt.Println("Diakhiri dengan dirun? : ", strings.HasSuffix(paragraf, "dirun"))

	// misahin string berdasarkan delimiter
	buah := strings.Split("salak, jambu, pisang, rambutan", ",")
	fmt.Println("Hasil pemisahan : ", buah)
	
	// replace string
	newParagraph := strings.ReplaceAll(paragraf, "app", "aplikasi")
	fmt.Println("Paragraf baru : ", newParagraph)


	// KONVERSI

	var a int = 88
	var b float64 = float64(a)
	var textA string = strconv.Itoa(a);

	fmt.Println("Nilai a : ", a)
	fmt.Println("Nilai b : ", b)
	fmt.Println("Teks a : ", textA)


	// PARSING
	// var text string = "100B"
	number, err := strconv.Atoi(textA)
	if err != nil {
		fmt.Println("Error : ", err.Error(), " Masukin angka yang bener nder")
		} else {
			fmt.Println("Angka : ", number)
		}

	// konversi & pasing boolean string
	// boolean ke string
	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println("Boolean ke string : ",str)
	
	// string ke boolean
	val,_ := strconv.ParseBool(str)
	fmt.Println("String ke boolean : ",val)

	// belajar variabel
	
	

	}
