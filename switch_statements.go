/* Switch Case */
/* Golang also supports a switch statement similar to that found in other languages such as, Php or Java.
Switch statements are an alternative way to express lengthy if else comparisons into more readable code based on the state of a variable. */

/* Golang - switch Statement */
/* The switch statement is used to select one of many blocks of code to be executed. */
package main

import "fmt"

func calculator() {
	var op string
	var num1, num2 float64
	fmt.Println("--- Initilizating Calculator ---")

	fmt.Print("Enter operator: +, -, *, / = ")
	fmt.Scanln(&op)

	fmt.Print("Please enter the first number = ")
	fmt.Scanln(&num1)

	fmt.Print("Please enter the second number = ")
	fmt.Scanln(&num2)

	switch op {
	case "+":
		fmt.Println("Output =", num1+num2)
	case "-":
		fmt.Println("Output =", num1-num2)
	case "*":
		fmt.Println("Output =", num1*num2)
	case "/":
		fmt.Println("Output =", num1/num2)
	default:
		fmt.Println("Invlid Input")

	}

}
