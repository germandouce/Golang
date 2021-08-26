package main

import (
	"fmt"
)
//Como decia...
//como todas las defs de la mujer son codigo repetido voy 
//a hacer q las herede del hombre--> Ver interfaces 2


type humano interface{
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface{
	respirar()
	comer()
	EsCarnivo() bool

}

type vegetal interface{
	ClasificacionVegetal() string
}


/*genero humano*/
type hombre struct{
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	esHombre bool //xa distinguir mujer de hombre xq esta 
					//ultima hereda el primero
}

//la mujer hereda todas las definiciones del hombre 
type mujer struct {
	hombre
}

//para que mi hombre y mi mujer sean de tipo humano...

//cada vez q yo llame a la funcion respirar le va a setear un true
//a la propiedad respirando del hombre
func (h* hombre) respirar() {h.respirando = true} 
func (h* hombre) comer() {h.comiendo = true} 
func (h* hombre) pensar() {h.pensando = true} 
func (h* hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre" } else {
		return "Mujer"
		}
}
  
//Ahora el hombre implementa la estructura humano


//Creo una funcion que en lugar e recibir un hombre reciba
//una intefaz humano
func HumanosRespirando(hu humano){
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando\n", hu.sexo())
}


//La interfaz me permite normalizar objetos que no tienen una union
//entre si y los permite clasificar dentro de un mismo tipo

//Entonces tanto el hombre como la mujer los hago pertenecer al tipo humano
func main(){
	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanosRespirando(Pedro)

	Maria :=new(hombre)
	//Maria.esHombre = false //no es necesario xq el objeto 
	//es x deafult false
	HumanosRespirando(Maria)
	
}

//se comporta de la misma manera q antes todo 0k!
