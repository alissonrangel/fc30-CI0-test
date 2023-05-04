package main

import "fmt"

func main() {
	fmt.Println(Soma(112, 10))
	fmt.Println("Soma realizada com sucesso")
}

func Soma(a int, b int) int {
	return a * b
}
