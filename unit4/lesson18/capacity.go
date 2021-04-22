package main

import "fmt"

func main() {
	beans := make([]string, 0, 0)
	var capacity int
	for i := 0; i < 10000000; i++ {
		beans = append(beans, "bean")
		if capacity == cap(beans) {
			continue
		}
		fmt.Printf("New capacity: %v\n", cap(beans))
		capacity = cap(beans)
	}
}
