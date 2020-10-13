package main

import "fmt"

func main () {
	var nilai int
	fmt.Print("Masukkan Bilangan : ")
	fmt.Scanf("%d", &nilai)
	if nilai > 0 {
		fmt.Println(nilai, " adalah Bilangan Positif \n")
	}else{
		fmt.Println(nilai, " adalah Bilangan Negatif \n")
	}
}