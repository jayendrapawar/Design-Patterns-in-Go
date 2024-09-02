package main

import "fmt"

type SedanCar struct { // This app code 
	Name string
}

func main() { // This is Client
	sc := getCar()
	fmt.Println(sc)
}

func getCar() SedanCar { // This is factory
	return SedanCar{Name: "Honda City"}
}
