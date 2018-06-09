package main

import (
	"fmt"
	"strconv"
)

type User struct {
	edad   int
	nombre string
}
type SuperUser struct {
	User
}

type inter interface {
	permisos() int
	name() string
}

func (usuario User) nombrea() string {
	return usuario.nombre
}
func main() {
	var a User
	a.nombre = "franc"

	fmt.Println(a.nombrea())
	fmt.Println(User{nombre: "jhon"})
	fmt.Println(a)
	var x *int
	entero := 5
	x = &entero
	*x = 55
	fmt.Println(*x)
	var algo int
	algo1 := "algo"
	y := 5
	arreglo := [3]int{2, 5, 6}
	matriz := [2][3]int{{2, 5, 6}, {11, 5, 0}}
	//arreglo := []int{2,5,6}
	slice := arreglo[:2]
	variableInt, _ := strconv.Atoi("4")
	fmt.Printf("suma %v"+strconv.Itoa(y+variableInt), algo1)
	fmt.Println(arreglo)
	fmt.Println(slice)
	fmt.Println(matriz)
	fmt.Scanf("%d\n", &algo)
	fmt.Printf("ssss %d\n", algo)
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
