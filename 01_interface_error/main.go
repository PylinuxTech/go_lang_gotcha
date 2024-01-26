package main

import (
	"fmt"
	"log"
)

type CustomError struct{}

func (d *CustomError) Error() string {
	return "Custom Error"
}

func GetError() *CustomError {
	return nil
}

func main() {

	var err error
	fmt.Printf("%T,%v\n", err, err)
	err = GetError()
	if err != nil {
		log.Fatal("This part shouldn't get executed")
	}
	fmt.Println("Program successfully executed")

}
