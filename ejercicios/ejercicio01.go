package ejercicios

import (
	"strconv"
)


func Ejercicio01F(input string) (int, string) {
	inputNum, err := strconv.Atoi(input)

	if err != nil{
		return 0, "hubo un error"
	}

	if inputNum>100{
		return inputNum, "es mayor a cien"
	}else{
		return inputNum,"es menor a cien"
	}
}