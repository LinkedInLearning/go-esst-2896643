package main

import (
	"fmt"
        "math"
)

	
type aireVolume interface {
    aire() float64
    volume() float64
}

type cube struct {
    longue_a float64
}

type sphere struct {
    radius float64
}

func (c cube) aire() float64 {
    return c.longue_a * c.longue_a * 6
}

func (c cube) volume() float64 {
    return c.longue_a * c.longue_a * c.longue_a
}

func (s sphere) aire() float64 {
    return s.radius * s.radius * 4 * math.Pi
}

func (s sphere) volume() float64 {
    return s.radius * s.radius * 4 * math.Pi / 3
}

func mesure(av aireVolume) {
    fmt.Println(av)
    fmt.Println("L'aire est égale à:", av.aire())
    fmt.Println("Le volume est égale à:",av.volume())
}

func main() {
    c := cube{longue_a: 10}
    s := sphere{radius: 4}
    mesure(c)
    mesure(s)
}