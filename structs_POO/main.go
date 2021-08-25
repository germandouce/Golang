package main

import (
	"fmt"
	"time"
	
	us "./user"
)

//los paquetes tienen q estar en la carpeta en la q estoy
//parado

type pepe struct {
	us.Usuario
}

func main() {
	fmt.Println("hola")
	//No estoy creando un new user, no estoy creando un nuevo usuario
	//estoy creando un usuario un usuario llamdo pepe que es una
	//herencia de la estructura llamada usuario
	u := new(pepe) //un user de tipo pepe
	u.AltaUsario(1, "Pablo Tilotta", time.Now(), true)
	fmt.Println(u.Usuario)
	fmt.Println("hola")
}
