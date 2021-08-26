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

type perro struct{
	respirando bool
	comiendo bool
	carnivoro bool
}

func (p* perro) respirar() {p.respirando = true} 
func (p* perro) comer() {p.comiendo = true} 
func (p* perro) EsCarnivoro() bool {return p.carnivoro}

//Estas funciones de tipo animal van a funcionar xa cada objeto de 
//que implemnete las 3 funciones de arriba
//La inetrfaz me permite agrupar en un tipo determinado de comportamiento
//a varios objetos q no tienen relacion entre si, en este caso
//con la funcion AnimalesRespirar

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

func main(){
	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	AnimalesRespirar(Dogo)
	totalCarnivoros =+ AnimalesCarnivoros(Dogo) 	

	fmt.Printf("Total carnivoros: %d", totalCarnivoros)
}

//se comporta de la misma manera q antes todo 0k!
