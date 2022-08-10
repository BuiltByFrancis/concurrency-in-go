package main

import "fmt"

/*
	The purpose of this example is to show a race condition, the data value may update in 3 places causing different behaviours

	1) first, this would cause nothing to print.
	2) after the if, before the print. This would cause the console to display 1.
	3) after the if. This would cause the console to display 0.
*/

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
