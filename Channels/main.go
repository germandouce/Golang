package main

import(
	"fmt"
	"time"
)

//Channels son vias q permiten que una Go rutine envie
//hacia otras funciones hacia main o hacia otra Go rutine
//De esta forma cada ejecucion paralela que desarrolla en
//el procesador, pueda dialogar con otra enviando informacion

//Un canal es un espacio de memoria de dialogo entre varias 
//rutinas y cuando se aloje un valor en ese espacio de memoria
//la rutina q lo esta pidiendo se ejecute
//Ese espacio de memoria tiene q ser de un tipo de dato, el 
//canal transporta un tipo de dato

func bucle(canal1 chan time.Duration){

	inicio := time.Now()	//guardo la hora de inicio
	for i :=0; i < 10000000000; i++{
	
	}
	final := time.Now()	//guaro la hora de finalizacion
	
	//le asigno al canal1 el valor resultante de esta funcion
	canal1 <- final.Sub(inicio)	//devuelve la diferencia entre final e inicio
}

func main(){
	canal1 := make(chan time.Duration) //canal de tipo time.duration
	//este tipo de dato lo vamos a usar para saber la diferencia entre distintas horas
	//xa saber cuanto duro esa rutina

	go bucle(canal1)
	fmt.Println("Llegue hasta aqui")
	
	//hasta que esto no ocurre no se detiene la ejecucion del programa
	msg := <- canal1	//es como el awake o promise de node.js. 
	//guardo en la var msg el contenido del canal. No puedo imprimir directamente el canal
	fmt.Println(msg)
	fmt.Println(canal1)	//Ver q no lo puedo imprimir ( me da una direc de memoria)
}