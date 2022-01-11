package main

import "fmt"

func main() {
	nombre, apellido := "Jhonny", "Rodríguez"
	const edad = 29

	var a int
	var b int8

	a = 17775
	b = 115

	c := a + int(b)

	fmt.Println(nombre, apellido, edad)
	fmt.Printf("Hola, yo soy %s %s \n", nombre, apellido)
	fmt.Printf("nombre's type: %T \n", nombre)
	fmt.Printf("apellido's type: %T \n", apellido)
	fmt.Println(c)
	fmt.Printf("value of c: %T", c)
}
