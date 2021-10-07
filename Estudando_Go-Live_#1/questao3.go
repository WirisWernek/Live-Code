package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// opção a
	// array := []string{}
	// for i := 0; i < 100; i++ {
	// 	array = append(array, "A")
	// 	fmt.Println(array)
	// }
	// fmt.Println(len(array))

	// opção b
	// palavra := []byte("asSASA ddd dsjkdsjs dk")
	// fmt.Println("Tamanho: ", len(palavra))
	// fmt.Println("Caracteres: ", utf8.RuneCount(palavra))

	// opção c
	// palavra := []byte("asSASA ddd dsjkdsjs dk")
	// palavra = append([]byte(palavra), " abc"...)
	// palavra_convert := string(palavra)
	// fmt.Println(palavra_convert)
	// fmt.Println("Tamanho: ", len(palavra))
	// fmt.Println("Caracteres: ", utf8.RuneCount(palavra))

	// opção d
	strings := "foobar"
	i := len(strings)
	caracter := []rune{}
	for i > 0 {
		r, s := utf8.DecodeLastRuneInString(strings)
		caracter = append(caracter, r)
		i--
		strings = strings[:len(strings)-s]
	}
	palavra := string(caracter)
	fmt.Println(palavra)
}
