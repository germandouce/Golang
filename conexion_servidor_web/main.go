package main

import(
	"net/http"
)

func main(){
	//va a entrar en home cuando el servidor detecta q estoy en
	//el local host o en barra (la rama principal se mi servidor web)

	//voy a crear la ruta q cuando detecto q entre en barra me ejecuta home

	// x ejemplo cuando en mi pagina voy a login, se ejcuta la funcion login
	//http.HandleFunc("/login", login)

	//en mi caso cunado voy a localhost
	http.HandleFunc("/", home)

	//Necisto decirle a Go q se ponga escuchar un puerto determinado y cuando escucha
	//u obtiene lo q quiere realiza una operacion de
	
	//Escucha y una vez recibe una peticion en ese puerto ejecuta una accion
					//el puerto
	http.ListenAndServe(":3000", nil)
}
	
	//xa el envio de respuestas			xa recibir respuestas

	func home(w http.ResponseWriter, r *http.Request){
	//cuando se abar el local host voy a mostrar index.html
	http.ServeFile(w, r, "./index.html")
	
	//xa ver la pagina abrir pestana en navegador y escribir localhost:3000
}