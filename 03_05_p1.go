package main

import (
	"encoding/json"
	"io/ioutil"
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
	listeAppels := Appel{
		Nom: "Heureux",
		T: []Tel{Tel{Portable: true, Num: "06 00 00 00 00"},
			Tel{Portable: true, Num: "0700000000"},
			Tel{Portable: false, Num: "01-00-00-00-00"},
		},
	}
	file, _ := json.MarshalIndent(listeAppels, "", " ")
 
	_ = ioutil.WriteFile("listeAppels.json", file, 0644)
}