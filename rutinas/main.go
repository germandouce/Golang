package main

import (
	"fmt"
	"strings"
	"time"
)

//Go rutines = similar a las promesas y a los async awakes de node.js

func miNombreLento(nombre string){
	letras := strings.Split(nombre,"") //devuelve una lista con los chars (como py)
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)	//quiero que cada letra tarde 1 seg en mostrarla
		fmt.Println(letra)
	}
}

func main(){
	go miNombreLento("German") //asi ejecuto mi nombre lento pero de manera asincrona
	fmt.Println("Estoy aqui")
	var x string	//paralela// se ejecuta la escritura del nombre
	fmt.Scanln(&x)
	//cuando le doy enter corta la ejecucion de la funcion asincrona xq se termino
	//la ultima instruccion del programa


}