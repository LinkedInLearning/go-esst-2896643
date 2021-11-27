package main

import (
	"fmt"
)

func main() {
	var monInt interface{} = 10
	l, res := monInt.(int)
	if res {
		fmt.Println("Succès:", l, "Valeur bool:", res)
	}
	m, res := monInt.(float64)
	if res {
		fmt.Println(m)
	} else {
		fmt.Println("MonInt ne contient pas de float64", "Valeur bool:", res)
	}
	
	a := monInt.(int)
	fmt.Println("Affichage a:", a)

	b := monInt.(bool)
	fmt.Println("Affichage b:", b)
}