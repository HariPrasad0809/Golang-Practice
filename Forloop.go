package main

import "fmt"

func loop() {
	for i := 1; i <= 30; i++ {
		fmt.Println("happy new year in 2025 Day:", i)
	}
}

func main() {
	fmt.Println("hello world")
	loop()
}
