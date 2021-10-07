package main

import "fmt"

func main() {
	// opção a
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// opção b

	// i := 1
	// FOR:
	// 	{
	// 		fmt.Println(i)
	// 		i++
	// 		if i > 10 {
	// 			os.Exit(0)
	// 		}
	// 		goto FOR
	// 	}

	// opção c
	// matriz := [2][5]int{}
	// 	var num int
	// 	num = 0
	// 	for i := 0; i < 2; i++ {
	// 		for j := 0; j < 5; j++ {
	// 			matriz[i][j] = num
	// 			num++
	// 		}
	// 	}
	// 	fmt.Println(matriz)
}
