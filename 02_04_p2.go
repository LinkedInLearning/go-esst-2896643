package main

import (
	"fmt"
)

func main() {	
	tab := []string{"Go", "Python", "Java"}
	for indice, valeur := range tab {
		fmt.Println("indice", indice, "valeur", valeur)
	}
}