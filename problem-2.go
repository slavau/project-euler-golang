package main

import "fmt"

func main() {

	x, y := 0, 1
	sum := 0

	for {
		x, y = y, x+y

		if y >= 4000000 {
			break
		}

		if y%2 == 0 {
			sum += y
		}

	}
	fmt.Println(sum)
}
