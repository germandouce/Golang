package main

import "fmt"

var estado bool

func if_else(){
	estado = true
	if estado = false; estado == true {
		fmt.Println(estado)
	}else if estado == false{
		fmt.Println("Es falso")
	}
}

func un_switch(){

	switch numero := 5; numero {
	case 1:
		fmt.Println(1)	
	case 2:
		fmt.Println(2)	
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)	
	case 5:
		fmt.Println(5)	
	default:
		fmt.Println("Mayor a 5")
	}
}

// if else
func main() {
	//if_else()
	un_switch()
}	