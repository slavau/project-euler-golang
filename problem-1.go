package main

import "fmt"
import "math"

func main() {
	var sum float64 = 0.0
	for i := 1.0; i < 1000; i++ {
		if math.Mod(i, 3) == 0 || math.Mod(i, 5) == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
