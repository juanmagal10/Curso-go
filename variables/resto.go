package variables

import (
	"fmt"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time


func RestoVariables() {

	Nombre ="pedro"
	Estado=true
	Sueldo=1577.77
	Fecha=time.Now()
	
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}