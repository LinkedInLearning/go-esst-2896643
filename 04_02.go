package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Il faut préciser un argument")
		return
	}

	y, erreur := strconv.Atoi(arguments[1])
	if erreur != nil {
		fmt.Println(erreur)
		return
	}

	triple := func(s int) int {
		return s * 3
	}		

	fmt.Println("Le triple de", y, "est", triple(y))
	
	double := func(s int) int {
		return s + s
	}
	fmt.Println("Le double de", y, "est", double(y))
	d, s := doubleTriple(y)
	fmt.Println("Résultat de l'impression de la fonction doubleTriple:",d, s)
}

func doubleTriple(x int) (int, int) {
	return x + x, x * 3
}
