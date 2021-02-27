package main

import "fmt"

type Person struct {
	Name string
}

type Persons struct {
	Persons []*Person
}

func main() {
	// ps := make([]Person, 5)
	// fmt.Println()
	// ps[0].Name = "Mike"
	// ps[1].Name = "Jon"
	// ps[2].Name = "NINA"
	// ps[3].Name = "joe"
	// ps[4].Name = "nancy"

	// fmt.Println(ps)

	p1 := Person{"Mike"}
	p2 := Person{"Jon"}
	ps := Persons{}
	ps.Persons = append(ps.Persons, &p1, &p2)

	for _, p := range ps.Persons {
		fmt.Println(p)
	}

}
