package main

import ("fmt"
		"strconv"
)

//Existen 3 tipos de variables ppal//
//Numericas, strings, booleanos

//Declaracion
//Bool se inicializa en false,
//Locales: arrancan en minuscula
//Globales (Xa verla desde otros paquetes): Arrancan con mayuscula


var numero int
var texto string
var status bool
var status_2 bool = true

var num float32 
var num_1 float64
var unsigned int

func main(){

	mostrarStatus()

	var numero_2 int
	numero_3 := 4
	//numero_3 :=10 //No puedo declarar asi dos veces seguidoas
	numero_3 = 15 //Si pued hacer esto, asignar  de nuevo
	//print(numero_2)
	
	//En una sola linea
	numero_2, numero_3, texto_1 := 2, 6, "hola"

	fmt.Println(numero_2)
	fmt.Println(numero_3)
	fmt.Println(texto_1)

	//No se permite cambiar el tipo de una variable
	//al asignarle otro valor
	var moneda float32 = 0
	//Lo prox tira error
	// var numero_2 = moneda

	//casteos:

	numero_2 = int(moneda)
	fmt.Println("numero_2 casteado de f a int: ",numero_2)

	texto = fmt.Sprintf("%f", moneda)
	fmt.Println("moenda casteado de f a str: ",texto)

	texto = strconv.Itoa(int(moneda))	//Itoa solo acepta enetro generico
	
	fmt.Println("Usando strconv: ",texto)

}

func mostrarStatus(){
	fmt.Print(status)
}






