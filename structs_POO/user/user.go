package user

//Ojo no se puede repetir los nombres de los pckages

import "time"
//Usuario va en mayuscula xq sera un objeto global
//que quiero poder ver ver desde afuera

type Usuario struct {
	Id int
	Nombre string
	FechaAlta time.Time	//Todos los paquetes q vienen de afuera van en mayuscula!! (x ej Printf)
	Status bool
}

func (this *Usuario) AltaUsario(id int, nombre string, fechaalta time.Time, status bool){
	//como usuario es un puntero cada vez q lo referencio tengo q usar this
	//this se usa para hacer referencia al objeto en el q estoy parado
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}