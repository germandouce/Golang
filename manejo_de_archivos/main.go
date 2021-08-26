package main

import(
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
)

func main(){
	//leoArchivo()
	//leoArchivo2()
	//graboArchivo()
	graboArchivo2()

}

func leoArchivo(){
	//Automaticamente abre, saca el contenido y cierra el buffer
	archivo, err := ioutil.ReadFile("./prueba_0.txt")
	//graba el conetenido en archivo y si no hay error true en error
	//La variable archivo es una lista con todos los numeros ascii de cada
	//caracter del archivos
	//fmt.Println(archivo)

	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))

	}
}

func leoArchivo2(){
	archivo, err := os.Open("./prueba_0.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		//sacanner nos permite capturar x teclado
		//ahora guardo una iterable con los elementos del acrhivo
		//en la variable sacenner
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() { //devueleve True o False y lee hasta q es falase es decir EOF
			registro := scanner.Text()	//devuelve un str con el contenido de cada linea
			fmt.Printf("Linea> "+ registro + "\n")
		}
	}
	archivo.Close()
}

func graboArchivo(){
	//create es como un open('file', 'w') ==> necesita un close
	//si ya existe sobreescribe!!
	archivo, err := os.Create("./nuevoArchivo_1.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return //xa q no ejecute lo q se encuentra abajo
	}
	
	fmt.Fprintln(archivo, "Esta es una linea nueva") // si ya existe lo sobreescribe!
	archivo.Close() // x el Create
}

func graboArchivo2(){
	//create es como un open ==> necesita un close
	fileName := "./nuevoArchivo_1.txt"
	//va a agregar nuevas lineas al final, sera com un como open('file', 'a')
	//OJO! No es propia de GO!!``
	if Appendear(fileName, "\nEsta es una segunda linea") == false{
		fmt.Println("Error al agrega una linea")
	}
}

func Appendear(archivo string, texto string) bool{
								//write only	//como el 'a' de pyhton	// permisos de escritura y lectura (mejor xa linux)
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0121)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	//si no hubo error, escribo la linea (al final xq lo abri como O_APPEND)
	_, err = arch.WriteString(texto)
	if err != nil{
		fmt.Println("Hubo un error")
		return false
	}
	fmt.Println("Se grabo correctamente el archivo")
	return true
}