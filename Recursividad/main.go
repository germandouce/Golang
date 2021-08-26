package main

import "fmt"

func main(){
	exponencia(2)
}

func exponencia (numero int){
	//primero hay q construir la salida 
	//(condicion de corte)
	if numero > 100000000 {
		return
	}
	fmt.Println(numero)
	exponencia(numero *2)
	//esto crea una pila en stack de llamadas 
	//en memoria donde se va a ir ejecutando
	//cuando saale de la funcion con return la 
	//pila de llamadas desaperece de memoria
}
