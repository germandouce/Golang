package main

import ("fmt"
		//"io/ioutil"
		"os"
		"log"
)

//DEFER - PANIC & RECOVER

//--- DEFER ---
func ej_defer(){
	archivo := "prueba_0.txt"
	f, err := os.Open(archivo)
	//como el archivo no esta en la carpeta execpiones
	//seguramente de un error

	//las instrucciones que siguen depues de defer
	//en caso de haber un error, 
	//NO SE EJECUTAN SECUENCIALMENTE!!
	//Recien se ejecutan cuando salgo de la funcion
	//pero se ejecuta SIEMPRE. 
	//y solo se ejecuta una sola cosa, no puede ser una lista de instrucciones
	defer f.Close()	//xa no consumir memoria en caso de que el archivo no exista

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		os.Exit(1)	//cierra la aplicacion o el programa directamente
	}

}

//PANIC & RECOVER
func ej_panic_recover(){
	//aseguro q lo q viene despues del defer se ejecuta si o si
	defer func(){	//asi le puedo dar varias instrucciones
		reco := recover()	//Solo si encuentro un panic guardo el resultado del ultimo panic
		if reco != nil{		//si mo hubo panic, se guarda un nil
			log.Fatalf("Ocurrio un error q genero panic\n %v",reco) // %v = 'variant'
			// hace 2 cosas, da un texto por pantalla y agrega la fecha y hora y ademas ejecuta una salida de programa
		}		
	}() //los parentesis son xq es una funcion que no devuelve nada (la estoy invocando)
	a :=1
	if a==1{
		panic("se encontro el valor de 1")
		//despues del panic no se puede hacer nignun control
		//xq corta el programa completamente, a menos que yo
		//coloque un defer q salta en caso de erro y ejecuta
		//lo q le pido
	}

}

func main(){
	//ej_defer()
	//ej_panic_recover()

}

//DEFER:
//Estructura q se ejecuta si o si cuando se detecta que
//una funcion se va x un return o por un error o x un 
//fin de funcion