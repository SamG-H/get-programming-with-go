package main

import "fmt"

func main() {
	str := "0"
	var strBool bool
	switch str {
		case "true", "yes", "1":
		strBool = true
		case "false", "no", "0":
		strBool = false
	}

	fmt.Printf("%v evaluates to %v\n", str, strBool)
}
