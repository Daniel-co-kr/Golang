package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusMany(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

//anonymous function
func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

//struct
type person struct {
	name string 
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 23
	return &p // related to garbage collector
}

 func main() {
	res := plus(1, 2)
	fmt.Println(res)

	res := plusMany(1, 2, 3)
	fmt.Println(res)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	nextInt := intSeq()

	fmt.Println(nextInt())

	// anonymous struct type
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true, 
	}

	fmt.Println(dog) // Rex true
 }