package main

import "fmt"

//solo existen for
//No existen los while, while do etc

func ciclo_1(){
	for i:= 0; i <=10; i++ {
		fmt.Println(i)
		}	
}

func ciclo_con_break(){
	for{
		numero := 0
		fmt.Println("Continuo")
		fmt.Println("ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100{
			break
		}
	}
}

func ciclo_con_continue(){
	var i = 0
	for i <10{
		fmt.Printf("\n valor de i: %d", i) //printf permite mascaras y verbos
		if i ==5{
			fmt.Printf("	Multiplicamos X 2\n")
			i = i*1
			continue
		}
		fmt.Printf("	Paso x aqui \n")
		i++
	}	
}	

func ciclo_con_go_to(){
	var i int = 0

	RUTINA:
		for i <10{
			if i ==4 {
				i = i+2
				fmt.Println("\n voy a RUTINA sumando 2 a i")
				goto RUTINA
			}
			fmt.Printf("Valor de i: %d \n", i)
			i++
		}

}

func main(){
	//ciclo_1()
	//ciclo_con_break()
	//ciclo_con_continue()
	ciclo_con_go_to()

}