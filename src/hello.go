package main

import "fmt"

func main() {
	for i := 0; i <= 30; i++ {
		fmt.Println(Fab(i))
	}
}

func Fab(index int) int {
	var fab = fabAdder()
	result := 0
	for i := 0; i <= index; i++ {
		result = fab(i)
	}
	return result
}

func fabAdder() func(index int) int {
	var result, last1, last2 int
	return func(index int) int {
		if index == 0 {
			result = 1
			last1 = 0
			last2 = 0
		} else if index == 1 {
			result = 1
			last1 = 1
			last2 = 1
		} else {
			result = last1 + last2
			last1 = last2
			last2 = result
		}
		return result
	}
}
