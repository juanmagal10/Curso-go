package main

import (
	"fmt"
	// "runtime"
	// "Go-desde-cero/variables"
	"Go-desde-cero/ejercicios"
)

func main() {
	// if os := runtime.GOOS; os=="Linux."||os=="OS X."{
	// 	fmt.Println("esto no es window")
	// }else{
	// 	fmt.Println("esto es ", os)
	// }

	numero, texto := ejercicios.Ejercicio01F("500")
	fmt.Println(numero);
	fmt.Println(texto);
}

