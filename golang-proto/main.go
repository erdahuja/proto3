package main

import (
	"fmt"

	"./src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := personpb.Person{
		Age:       12,
		FirstName: "Deepak",
		LastName:  "Ahuja",
		IsVerfied: true,
	}
	fmt.Println(sm)
	sm.FirstName = "Renamed"
	fmt.Println(sm)
	fmt.Println(sm.Height)
}
