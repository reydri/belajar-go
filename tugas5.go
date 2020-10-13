package main

import "fmt"

func main () {
	var nilai int
	fmt.Print("Masukkan Bilangan : ")
	fmt.Scanf("%d", &nilai)
	for i := 0; i < nilai; i++ {
		for j := i; j < nilai; j++ {
			fmt.Print(j, "")
		}
		fmt.Println()
	}
}