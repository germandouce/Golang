package main

import "fmt"

var tabla[10] int
//		[fil][col]
var matriz [5][7] int
func main(){
	//ARRAYS
	// tabla[0]=1
	// tabla[5]=15
	// tabla_2 := [10] int{5,6,7,8,}
	// fmt.Println(tabla_2)
	// fmt.Println(tabla)

	// for i := 0; i < len (tabla_2); i++{
	// 	fmt.Println(tabla_2[i])
	// }

	//MATRICES
	// matriz[3][5]=1
	// fmt.Println(matriz)

	//-------SLICES----: 
	//Vectores dinamicos, es decir que puedo ampliar
	//la dimension en tiempo de ejecucion ("Parecidos a listas")
	
	//Creacion --> slice := []type
	//Ejemplo con un append 
	// matriz:= []int{2,5,4}
	// fmt.Println(matriz)
	// matriz = append(matriz, 8)
	// fmt.Println(matriz)

	//variante2
	//variante2()
	//variante3()
	//variante4()
}

func variante2(){
	elementos := [5]int{1,2,3,4,5}
	porcion :=	elementos[3:]	//sin incluir el 3
	porcion_2 := elementos[2:4] //sin el 2 inclusive el 4

	fmt.Println(porcion)
	fmt.Println(porcion_2)
}

func variante3(){
	//Adicionar elememtos a un Slice
	//comando make crea el slice con un largo(inicial), capacidad(en memoria)
	elementos := make([]int,5,20)
	//si itero, itero sobre 5
	//len -> longitud, elementos relaes
	//cap -> capacity, max cant de elementos
	
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func variante4(){
	//Extiendo un slice --> append
	//Ojo! aca a diferencia de Python append es 
	//una funcion no un metodo!!
	//append(slice,ele_a_cargar)
	nums := make ([]int,0,0)
	for i := 0; i < 100; i++{
		nums=append(nums, i)		
	}
	//Go aumenta la capacidad automaticamente, en este caso en 8 bytes
	//asi: 2, 4, 6, 8, 32, 64, 128, 256, 512, etc 2^p
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))

}