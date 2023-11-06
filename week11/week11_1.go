package main

import "fmt"

func main() {
	primes := [3]int{2, 3, 5} //initalize by array literal
	primes[2] = 6

	for i := 0; i < 3; i++ {
		fmt.Println(primes[i])
	}
	test := [5]bool{true, true, true}
	fmt.Println(test[3])
	fmt.Println(test)

	fmt.Println("%v\n", primes)
	fmt.Println("%v\n", test)d
}
