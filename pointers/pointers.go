package main

import "fmt"

func main() {
	i, j := 23, 2704

	p := &i
	fmt.Println(*p)

	*p = 34
	fmt.Println(i)

	p = &j
	*p = *p / 37

	fmt.Println(j)
}
