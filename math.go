package main

import "fmt"

func main() {
	fmt.Println(soma(112, 10))
	fmt.Println("Soma realizada com sucesso")
}

func soma(a int, b int) int {
	return a * b
}
