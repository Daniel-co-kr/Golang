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

//defer, panic, recover
func myPanic() {
	panic("a problem")
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

	//defer
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("main")

	//panic -> 현재 함수 실행 즉시 중단 & defer 함수들을 LIFO 방식으로 실행
	//recover -> 반드시 defer 함수 내에서 사용
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error: ", r)
		}
	}()
	
	myPanic()
	fmt.Println("After mayPanic()") //실행 x 

	//channel
	messages := make(chan string)

	go func() { messages <- "ping" }() //송신

	msg := <- messages //수신 -> 데이터를 받을 때까지 대기(blocking 상태)
	fmt.Println(msg)
	//goroutine이 진행되는 동안 main goroutine은 병렬적으로 진행 가능
	//main goroutine이 종료 시 프로그램도 종료
 }