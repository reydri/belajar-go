package main

import "fmt"

func main () {
	var nilai int
	fmt.Print("Masukkan Bilangan : ")
	fmt.Scanf("%d", &nilai)
	if nilai%2==0 {
		fmt.Print(nilai, " adalah Bilangan Genap \n")
	}else{
		fmt.Print(nilai, " adalah Bilangan Ganjil \n")
	}
}