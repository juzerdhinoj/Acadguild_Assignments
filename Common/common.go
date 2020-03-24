package Common

import (
	"fmt"
	"reflect"
)

var Num1 = 1
var Num2, Num3 int = 2, 3

var name string = "Juzer"

func Variable_Types() {

	Num4 := 4.2

	fmt.Println("Variable Types")

	fmt.Println(Num1, "-", reflect.TypeOf(Num1).String())
	fmt.Println(Num2, "-", reflect.TypeOf(Num2).String())
	fmt.Println(Num3, "-", reflect.TypeOf(Num3).String())
	fmt.Println(Num4, "-", reflect.TypeOf(Num4).String())
	fmt.Println(name, "-", reflect.TypeOf(name).String())

}
