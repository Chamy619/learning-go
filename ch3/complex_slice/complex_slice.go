package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4) // {1,2,3,4}
	y := x[:2:2]              // {1,2}
	z := x[2:4:4]             // {3,4}
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50) // x: {1,2,30,40,50} y: {1,2,30,40,50}
	x = append(x, 60)         // x: {1,2,30,40,60} z: {30,40,60}
	z = append(z, 70)         // x: {1,2,30,40,70} z: {30,40,70}
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
