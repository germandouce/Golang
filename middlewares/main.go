package main

import "fmt"

//middlewares: interceptores
//Son como munecas rusas una adentro de la otra que nos 
//permiten encapsular la ejecucion de una funcion ya 
//existente con otra funcion nueva.
//Ejecuta cierto codigo previamente al llamado de la 
//funcion que ya tenemos grabada.
//(mio: es una especie de trigger before de sql)

/*Def profe: Son interceptores que permiten ejecutar intrucciones 
comunes a varias funciones que reciben y devuelven los mismos 
tipos de variables */

//si creo un interceptor que opera entre varias funciones esas
//funciones tienen q ser iguales en cuanto al tipo de valor q 
//reciben y devuelven

func sumar(a,b int) int {
	return a+b
}

func restar(a,b int) int {
	return a-b
}

func multiplicar(a,b int) int {
	return a*b
}

//recibe una funcion como parametro
					//recibe				//devuelve
func operacionesMidd( f func(int, int) int ) func(int, int) int {
	//el middleware funciona como un pasamano,
	//le paso la referencia a mi funcion la recibe,
	//la procesa y devuelve la misma funcion para 
	//q luego de eso ejecute lo q esta en la funcion
	return func (a, b int) int {
		fmt.Println("Inicio de operacion")
		/*aqui se hacen los llamados a operaciones de 
		encriptacion seguridad etc antes del return*/
		return f(a,b)	
	}
}

var result int
func main(){
	//Si quiseses hacer las 3 operaciones sumar, restar y multiplicar
	//podria agregarles un fmt.p.. a cada una pero es mas facil...
	
	fmt.Println("inicio")

	//la funcion middleware
	//voy a grabar en result el resultado de la operacion
	//midd evaluada en la funcion con la que voy a tener q trabajar
	// y cuales son los avlores de entrada q va a tener esa funcion

	//encapsulo las 3 funciones en un middleware
	result = operacionesMidd(sumar)(2,2)
	fmt.Println(result)
	
	result = operacionesMidd(restar)(10,6)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(1,4)
	fmt.Println(result)

}