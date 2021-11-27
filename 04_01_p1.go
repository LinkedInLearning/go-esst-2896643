package main

import (
	"fmt"
)

func main() {
	m:=func(){
		fmt.Println("Fonction anonyme")
	}
	m()
}