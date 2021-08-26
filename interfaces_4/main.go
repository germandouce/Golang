package main

import (
	"fmt"
)
//Como decia...
//como todas las defs de la mujer son codigo repetido voy 
//a hacer q las herede del hombre--> Ver interfaces 2


type SerVivo interface{
	estaVivo() bool
}

type humano interface{
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool

}

type vegetal interface{
	ClasificacionVegetal() string
}
//La interfaz me permite normalizar objetos que no tienen una union
//entre si y los permite clasificar dentro de un mismo tipo

//Entonces tanto el hombre como la mujer los hago pertenecer al tipo humano

/*Genero Animal*/
type perro struct{
	respirando bool
	comiendo bool
	carnivoro bool
	vivo bool
}

func (p* perro) respirar() {p.respirando = true} 
func (p* perro) comer() {p.comiendo = true} 
func (p* perro) EsCarnivoro() bool {return p.carnivoro}
func (p* perro) estaVivo() bool {return p.vivo}

//Estas funciones de tipo animal van a funcionar xa cada objeto de 
//que implemnete las 3 funciones de arriba
//La inetrfaz me permite agrupar en un tipo determinado de comportamiento
//a varios objetos q no tienen relacion entre si, en este caso
//con la funcion AnimalesRespirar

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
	vivo bool
}

func (h* hombre) respirar() {h.respirando = true} 
func (h* hombre) comer() {h.comiendo = true} 
func (h* hombre) pensar() {h.pensando = true} 
func (h* hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre" } else {
		return "Mujer"
		}
}


func (h * hombre) estaVivo() bool {return h.vivo}

func HumanosRespirando(hu humano){
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando\n", hu.sexo())
}

func AnimalesRespirar(an animal){
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

/*serVivo */
func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}

func main(){
	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros =+ AnimalesCarnivoros(Dogo) 	

	fmt.Printf("Total carnivoros: %d \n", totalCarnivoros)

	fmt.Printf("Total vivo = %t", estoyVivo(Dogo))

}

//Un objeto esta implementando distintos tipos de interfaz