/* Constants */
/* A constant is a name or an identifier for a fixed value. The value of a variable can vary.
but the value of a constant must remain constant. */

package main

import "fmt"

/* Declaring (Creating) Constants */
/* The keyword const is used for declaring constants followed by the desired name and the type of value the constant will hold.
You must assign a value at the time of the constant declaration, you can't assign a value later as with variables. */

func dec() {
	fmt.Println(PRODUCT)
	fmt.Println(PRICE)
}

/* Multilple Constants Declaration Block */
/* Constants declaration can to be grouped together into blocks for greater readability and code quality. */

const (
	PRODUCT  = "Mobile"
	QUANTITY = 50
	PRICE    = 50.50
	STOCK    = true
)

func multicons() {
	fmt.Println(QUANTITY)
	fmt.Println(PRICE)
	fmt.Println(PRODUCT)
	fmt.Println(STOCK)
}
