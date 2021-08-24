package main

import (
	"fmt"
	"os"
	"bufio"
	)	
// 	"os"
// 	"buffio"
// )

func main(){

	var numero_1 int
	var numero_2 int

	var resultado int
	var leyenda string

	fmt.Println("Ingrese numero 1: ")
	fmt.Scanln(&numero_1)
	//OJO BUG:
	// linux: \n
	//windows: \n \r


	fmt.Println("Ingrese numero 2:")
	fmt.Scanln(&numero_2)

	//fmt.Printf("total: %d", numero_1 + numero_2)

	fmt.Println("Descripcion: ")
	
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		leyenda = scanner.Text()
	}
	resultado = numero_1 + numero_2
	fmt.Println(leyenda, resultado)

}