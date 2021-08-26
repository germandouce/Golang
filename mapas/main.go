package main

import "fmt"

func main(){
	//forma 1: con make
	//paises := make(map[string]string)
	//fmt.Println(paises)

	// paises["Islas_fiji"] = "Suva"
	// paises["Angola"] = "Luanda"

	// fmt.Println(paises)

	//forma 2
	//cant elementos. Reservo espacio xa 5 paises
	//despues puedo agregarle elementos
	//paises := make(map[string]string, 5)

	//forma 3: con asignacion directa
	campeonato := map[string]int{
		"Barcelona": 39,
		"Real Madrid": 38,
		"Chivas": 37,
		"Boca Juniors": 1000000}
	
	fmt.Println(campeonato)
	//Automaticamente ordena alfabeticamente x clave
	
	//Agregar elementos al mapa
	campeonato["Ferro"]=25
	//fmt.Println(campeonato)

	//cambio un valor
	campeonato["Chivas"]= 57
	//fmt.Println(campeonato)

	//Borrar un elemento del mapa
	delete(campeonato, "Real Madrid")
	//fmt.Println(campeonato)

	//Recorrer el mapa
	//El range ya no ordena alfabeticamente
	for equipo, puntaje := range campeonato{
		fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	}

	//Chequer si un equipo existe en mi mapa
	puntaje, existe := campeonato["Mineiro"]
	//Una clave no existente es el int 0
	//y su valor es el booleano false
	fmt.Printf("El Puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}	