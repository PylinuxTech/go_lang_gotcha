package main

import "fmt"

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main() {

	d1 := data{"value"}
	d1.print()

	var in printer = &d1 // error
	fmt.Println(in)

	m := map[string]*data{"key": &data{"value"}}
	m["key"].print()

}
