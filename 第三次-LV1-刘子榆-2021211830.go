	package main

	import (
		"fmt"
		"time"
	)

	func main() {
		go repeat1("A", 10)
		go repeat2("B", 10)
		go repeat3("C", 10)
		time.Sleep(time.Millisecond*10000)
	}
	func repeat1(str string, count int) {
		for i := 0; i < count; i++ {
			time.Sleep(time.Millisecond*1)
			fmt.Println(str)
			time.Sleep(time.Millisecond*1000)
		}
	}
	func repeat2(str string, count int) {
		for i := 0; i < count; i++ {
			time.Sleep(time.Millisecond*5)
			fmt.Println(str)
			time.Sleep(time.Millisecond*1010)
		}
	}
	func repeat3(str string, count int) {
		for i := 0; i < count; i++ {
			time.Sleep(time.Millisecond*7)
			fmt.Println(str)
			time.Sleep(time.Millisecond*1020)
		}
	}


