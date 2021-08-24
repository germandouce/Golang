package main

import "fmt"

func main() {
	//fmt.Println(uno(5))

	// numero, estado := dos(2)

	// fmt.Println(numero)
	// fmt.Println(estado)

	//fmt.Println(uno_bis(5)) //tambien muestra 10

	fmt.Println(calculo(5,46))
	fmt.Println(calculo(2,23,45,68))
	fmt.Println(calculo(5))
	fmt.Println(calculo(5,46,17,25,26,98,17))
}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func uno_bis(numero int) (z int) {
	z = numero *2
	return
}

//parametros variables
func calculo(numero ...int) int {
	total := 0
	//range es una instruccion q devuelve 2 datos
	//el primero es el numero de elemento
	//el segundo, es el elemento en si
	for _, num := range numero{	//_ es xq no necesito ni uso el valor
		total = total + num
	}
	return total

}

func calculo_bis(numero ...int) int {
	total := 0
	//range es una instruccion q devuelve 2 datos
	//el primero es el numero de elemento
	//el segundo, es el elemento en si
	for item, num := range numero{	//_ es xq no necesito ni uso el valor
		total = total + num
		fmt.Printf("\n Item %d\n", item)
	}
	return total

}