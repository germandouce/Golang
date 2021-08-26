package main

import (
	"fmt"
)

//Con las interfaces podremos grabar 
//comportamientos de determinados 
//objetos

//En una interface se definen los metodos q 
//vamos a usar xa implementar esa interface

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
}

type mujer struct {
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
}

//para que mi hombre y mi mujer sean de tipo humano...

//cada vez q yo llame a la funcion respirar le va a setear un true
//a la propiedad respirando del hombre
func (h* hombre) respirar() {h.respirando = true} 
func (h* hombre) comer() {h.comiendo = true} 
func (h* hombre) pensar() {h.pensando = true} 
func (h* hombre) sexo() string {return "Hombre"} 
//Ahora el hombre implementa la estructura humano

//Idem xa mujer...
func (m* mujer) respirar() {m.respirando = true} 
func (m* mujer) comer() {m.comiendo = true} 
func (m* mujer) pensar() {m.pensando = true} 
func (m* mujer) sexo() string {return "mujer"} 

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
	HumanosRespirando(Pedro)

	Maria :=new(hombre)
	HumanosRespirando(Maria)
	
}

//como todas las defs de la mujer son codigo repetido voy 
//a hacer q las herede del hombre--> Ver interfaces 2
