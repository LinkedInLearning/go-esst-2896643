package main

import (
	"fmt"
	"time"
)

func main() {
	
	date_a := time.Now()
	fmt.Println(date_a)
	date := time.Date(2014, 10, 9, 04, 20, 58, 548626532, time.UTC)
        fmt.Println(date)
	fmt.Println("Année",date.Year(),"Mois",date.Month(),"Jour",date.Day(),"Heure",date.Hour())
	fmt.Println("Minutes",date.Minute(),"Secondes",date.Second(),"Nanosecondes",date.Nanosecond(),"Fuseau horaire",date.Location())

	dur := date_a.Sub(date)
    	fmt.Println(dur)

	fmt.Println("Durée en heures",dur.Hours())

	fmt.Println(date_a.Add(-dur))
  	
}

