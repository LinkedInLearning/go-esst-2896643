package main

import (
	"fmt"
)

func main() {
	ville_1 := ville{}
	ville_1.nom = "Montréal"
	ville_1.superficie = 431.5
	ville_2 := ville{ "Rabat" , 117 }
	fmt.Println( "nom : " , ville_1.nom)
	fmt.Println( "superficie : " , ville_1.superficie)
	fmt.Println( "nom : " , ville_2.nom)
	fmt.Println( "superficie : " , ville_2.superficie)
}

type ville struct {
	nom string
	superficie float32
}