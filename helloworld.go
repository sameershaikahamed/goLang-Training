package main

import (
	"fmt"
	"reflect"
)

const num = 100

func main() {

	println("Hello World")

	println("Test, Test1 ,Test2 ") //double quotes for the string

	fmt.Println("TEST")
	//Identifiers --user defined variable
	fruit := "apple"

	fmt.Println("fruit")
	fmt.Println(fruit)

	fmt.Println("fruit", fruit) //fruit apple

	fmt.Println("Address location of fruit", &fruit) //Address location of fruit 0xc000086050

	fmt.Println(fruit == fruit) //true

	var a = 10
	var b = 20
	var c = a + b
	fmt.Println("Hello World", c)

	//fmt.Println(' a= ',a)

	//myFunc()

	myFunc(c)

}
func myFunc(param int) {

	fmt.Println("Value passed to myFunc:", param)

	fmt.Println("Value passed to myFunc:", reflect.TypeOf(param)) //int

	fmt.Println("large Value passed to myFunc:", reflect.TypeOf(143243243434343)) //int

	var data float64 = 10.2
	fmt.Printf("data = %v, type = %T\n", data, data) // Correct usage of formatted output

	fmt.Println("Value passed to myFunc Const:", num)

	const num2 = 20

	fmt.Println("Value passed to myFunc Const:", num2)

	var blue int = 255 //overflow -127 to 128

	fmt.Printf("space= %T, type=%d Binary Value=%b\n", blue, blue, blue)

}

/*func myFunc() {

	fmt.Println("Second Func ")

	var num1, num3 int = 1, 3

	//exportable

	var num2 int = 2

	fmt.Println("Num1 :", num1)

	fmt.Println("Num3 :", num3)

	fmt.Println("Num2 :", num2)
}*/
