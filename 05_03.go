package main

import (
	"fmt"
)

type emploi struct {
    Titre string
    ID string
}

func (e emploi ) String() string {
    return fmt.Sprintf("L'emploi %q a comme identifiant %s.", e.Titre, e.ID)
}

func main() {
    e := emploi{
        Titre: "Développeur GO",
        ID: "G00201",
    }
    fmt.Println(e.String())
}