// displayVariablesAndTypes example

package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	color      string
	horsepower int
}

func main() {
	v1 := 5                 //int
	v2 := 356.58461235649   //float64
	v3 := "Hello World!"    //string
	v4 := Car{}             //Car
	v5 := Car{color: "red"} //Car
	v6 := true              //bool

	displayVariablesAndTypes(v1, v2, v3, v4, v5, v6)
}

func displayVariablesAndTypes(elements ...interface{}) {
	for index, element := range elements {
		fmt.Printf("Variable #%d: %v (%s)\n", index+1, element, reflect.TypeOf(element))
	}
}
