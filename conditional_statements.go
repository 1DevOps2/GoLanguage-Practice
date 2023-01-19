/* If Else - Golang Conditional Statements
Like most programming languages, Golang borrows several of its control flow syntax from the C-family of languages. In Golang we have the following conditional statements:

- The if statement - executes some code if one condition is true
- The if...else statement - executes some code if a condition is true and another code if that condition is false
- The if...else if....else statement - executes different codes for more than two conditions
- The switch...case statement - selects one of many blocks of code to be executed
*/

/* Golang - if Statement */
/* The if statement is used to execute a block of code only if the specified condition evaluates to true.
Syntax:
		if  condition {
	// code to be executed if condition is true
		}
*/
package main

import "fmt"

func ifstatement() {
	var lang = "Go-Language"
	var x bool = true
	if x {
		fmt.Println(lang)
	}

}

/* Golang - if-else Statement */
/* The if....else statement allows you to execute one block of code if the specified condition is evaluates to true and another block of code if it is evaluates to false.

Syntax:
		if  condition {
			// code to be executed if condition is true
		} else {
			// code to be executed if condition is false
		}
*/

func ifelse() {
	city := "Multan"
	if city == "Multan" {
		fmt.Println("Multan Belongs to Pakistan")
	} else {
		fmt.Println("This City is not Belong to Pakistan")

	}
}

/* Golang - if...else if...else Statement */
/* The if...else if...else statement allows to combine multiple if...else statements.

Syntax:
		if  condition-1 {
			// code to be executed if condition-1 is true
		} else if condition-2 {
			// code to be executed if condition-2 is true
		} else {
			// code to be executed if both condition1 and condition2 are false
		}
*/

func ifelseif() {
	city := "Chaman"
	if city == "Multan" {
		fmt.Println("This city belongs to punjab")
	} else if city == "Karachi" {
		fmt.Println("This city belongs to Sindh")
	} else if city == "Chaman" {
		fmt.Println("This city belongs to Balochistan")
	} else {
		fmt.Println("This City is not belongs to Pakistan")
	}
}

/* Golang - if statement initialization */
/* The if statement supports a composite syntax where the tested expression is preceded by an initialization statement.

Syntax:
		if  var declaration;  condition {
			// code to be executed if condition is true
		}
*/

func ifinit() {
	if x := 100; x <= 100 {
		fmt.Println("Valid Output")
	}
	if x := true; x {
		fmt.Println("Boolean Value is True")
	}
}
