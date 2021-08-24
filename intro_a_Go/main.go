package main
//Cada archivo en Go debe comenzar con la palabra reservada package
//El nombre del archivo y el paquete tienen q ser iguales!
//La idea de los paquetes es q ue depues se puedan importar en otros archivos

//"xa un solo paquete"
//("xa varios paquetes")
import "fmt"	//xa mostrar strings x pantalla

//Ojo! Go no nos deja guardar el archivo con malas practicas!
//Borra lo q no se usa!

//En Go todo son funciones
//Siempre comineza x la funcion main()
//Es sensible a la identacion

func main(){
	fmt.Println("Hola mundo!")
}