package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func Ejercicio2(){
	fmt.Println(("ingresa el numero"))
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan(){
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil{
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}
	for i:=0;i<11;i++{
		fmt.Println(numero*i)
	}
}