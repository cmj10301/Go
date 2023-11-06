package main

import "fmt"

func main() {
	// 포인터 : 메모리 주소를 값으로 가지는 타입. & 주소연산자,
	var a int = 10
	var pa *int
	pa = &a

	fmt.Println(&a, a)
	fmt.Println(pa)
	fmt.Println(&pa)
	fmt.Printf("%T %T\n", a, pa)
}
