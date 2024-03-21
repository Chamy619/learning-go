package main

import "fmt"

func main() {
	x := 10
	x = 20
	fmt.Println(x)
	x = 30 // go vet, go compiler 는 잡아내지 못하지만 golangci-lint 는 잡아냄
}
