package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Appel struct {
	Nom string
	T []Tel
}

type Tel struct {
	Portable bool
	Num string
}

func main() {
	var d Appel
	file, _ := ioutil.ReadFile("listeAppels.json")
	json.Unmarshal(file, &d)
	fmt.Println( "Données fichier json : " , d)
}