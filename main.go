package main

import (
	"GoLang_Practice/week1/Common"
	"fmt"
)

func main() {
	fmt.Println("In Main")
	Common.Variable_Types()

	calculate()
}

func calculate() {

	add := Common.Num1 + Common.Num2

	fmt.Println("Addition: ", add)

	multiply := Common.Num3 * Common.Num3

	fmt.Println("Multiplication: ", multiply)

}
