package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	errs = "There's an error"
)

/* Declaration with initialization */
func demo() {
	var x float64 = 10.23
	var y int
	fmt.Println(x)
	fmt.Println(y)

}

/* Declaration without initialization*/
func builtsum() {
	var x int64
	var y int64
	var res int64
	fmt.Println("--- Built-in Sum Calculator ---")
	x = 10
	y = 6
	res = x + y
	fmt.Println("output is =", res)

}

/* Multiplication Calculator */
func multiply() {
	var x int64
	var y int64
	var output int64
	fmt.Println("--- Multiplication Calculator ---")

	fmt.Print("Please enter the first number = ")
	fmt.Scan(&x)

	fmt.Print("Please enter the second number = ")
	fmt.Scan(&y)

	output = x * y
	fmt.Println("output is =", output)

}

/* Declaration with type inference --- omit the variable type from the declaration */

func inference() {
	year := 2023
	var country = "pakistan"
	y := "anaskhan"
	fmt.Println("year datatype = ", year, reflect.TypeOf(year))
	fmt.Println(reflect.TypeOf(country))
	fmt.Printf("y data type is %T", y)

}

/* Convert string to integer type */
func convert() {
	strVar := "10000"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	strVar1 := "2000"
	intVar1, err := strconv.ParseInt(strVar1, 0, 65)
	fmt.Println(intVar1, err, reflect.TypeOf(intVar1))

}

/* convert Boolean Type to String */
func booltostr() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))

	var s string = strconv.FormatBool(true)
	fmt.Println(s, reflect.TypeOf(s))
}

/* Declaration of multiple variables */
/* Golang allows you to assign values to multiple variables in one line. */

func multivar() {
	var fname, lname = "anas", "khan"
	m, n, o := 1, 2, 3
	reg, batch := 02, "BSSE"

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(reg, "-", batch)

}

/* Short Variable Declaration in Golang */
/* The := short variable assignment operator indicates that short variable declaration is being used.
There is no need to use the var keyword or declare the variable type. */

func shortvar() {
	name := "Anas Khan"
	fmt.Println(reflect.TypeOf(name))
}

/* Scope of Golang Variables Defined by Brace Brackets */
/* Golang uses lexical scoping based on code blocks to determine the scope of variables.
Inner block can access its outer block defined variables, but outer block cannot access inner block defined variables.*/

var countryname = "Spain"

func bracevar() {
	fmt.Println(countryname)
	x := true

	if x {
		y := 1
		if x != false {
			println(y)
		}
	}

}

/* Zero Values */
/* If you declare a variable without assigning it a value, Golang will automatically bind a default value (or a zero-value)
to the variable. */

func zerovar() {
	var price float32
	fmt.Println(price)

	var ID int
	fmt.Println(ID)

	var orders string
	fmt.Println(orders)

	var trust bool
	fmt.Println(trust)
}

/* Golang Variable Declaration Block */
/* Variables declaration can be grouped together into blocks for greater readability and code quality. */
func blockvar() {
	var (
		product  = "Mobile"
		quantity = 50
		price    = 50.50
		inStock  = true
	)
	fmt.Println(quantity)
	fmt.Println(price)
	fmt.Println(product)
	fmt.Println(inStock)
}
