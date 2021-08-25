package main
import "fmt"

//En Go las funciones son tipos de datos!

var Calculo func(int, int) int = func (num1 int, num2 int) int{		
	return num1 + num2
}

func main(){
	fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5,7))
	
	//Restamos									//Similarmente a una sobrecarga
	Calculo = func(num1 int, num2 int) int {	//Puedo redefinir funciones!!
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 4 = %d \n", Calculo(6,4))

	
	//Dividimos									//Similarmente a una sobrecarga
	Calculo = func(num1 int, num2 int) int {	//Puedo redefinir funciones!!
		return num1/num2
	}
	fmt.Printf("Dividimos 12/3 = %d \n", Calculo(6,4))

	//Ver q calculo siempre se llama de la misma manera

	Operaciones()

	/*CLOSURES*/
	tablaDel :=2
	//Mi tabla ahroa es una funcion 
	//Toma como valor la funcion que devuelve Tabla
	Mitabla:= Tabla(tablaDel)	
	for i := 1; i < 11; i++{	
		fmt.Println(Mitabla())
	}
}

func Operaciones(){
	resultado := func() int{	//una variable de tipo funcion
	var a int = 23				//dentro de otra funcion
	var b int = 13
	return a + b
	}

	fmt.Println(resultado())
}

//CLOSURES

//Una funcino tabla q me devuelve otra funcion
func Tabla(valor int) func() int{
	numero := valor
	secuencia := 0 //secuencia esta "protegida" xq se encuentra solo dentro de tabla

	return func() int{
		secuencia +=1
		return numero*secuencia
	}
}