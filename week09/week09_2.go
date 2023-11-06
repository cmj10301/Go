package main

import "fmt"

// pass by pointer (call by pointer)
func swap(n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func main() {
	a := 10
	b := 20
	c := &a //integer 타입의 변수 주소는 integer 타입의 포인트 변수로만 받을 수 있다.

	fmt.Printf("%T\n", c)

	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}
