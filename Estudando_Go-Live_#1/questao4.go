package main

import "fmt"

func main() {
	slice := []float64{51.6, 58.78, 32, 14.98}
	var soma float64
	for i := 0; i < len(slice); i++ {
		soma = soma + slice[i]
	}
	x := len(slice)
	tam := float64(x)
	media := soma / tam

	fmt.Println("Média: ", media)

}
