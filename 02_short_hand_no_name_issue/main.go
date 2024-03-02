package main

import "fmt"

func main() {
	var name_variable int
	fmt.Println(name_variable)

	a := 12
	fmt.Printf("%T,%v\n", a, a)

	b, c := 12, 2.4
	fmt.Println(b, c)

	a, d := 12, 23
	fmt.Println(a, d)

	type CompositeType struct {
		a int
	}

	CompositeTypeVar := CompositeType{}

	fmt.Println(CompositeTypeVar)

	CompositeTypeVar.a = 45
	fmt.Println(CompositeTypeVar)

	var e int
	CompositeTypeVar.a, e = 23, 23
	fmt.Println(CompositeTypeVar, e)

	// array[i],new_variable := 23,23

}
