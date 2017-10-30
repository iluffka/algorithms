package main

import "fmt"

func main() {
	//fmt.Println(FibonacciRecursive(50))
	//fmt.Println(FibonacciNotRecursive(120))
	var seq []int
	fmt.Println(FillSequence(80, &seq))
	fmt.Println(seq)
}

func FibonacciRecursive(n int) int64 {
	if n == 0 || n == 1 {
		return 1
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func FibonacciNotRecursive(n int) int64 {
	if n == 0 {
		return 1
	}
	var prev = int64(1)    //F(0)
	var current = int64(1) //F(1)
	for i := 2; i <= n; i++ {
		temp := current
		current += prev
		prev = temp
	}
	return current
}

/*
*Рекурсивный поиск Фибаначи с запоминанием
 */
func FillSequence(index int, sequence *[]int) int {
	if index < 0 || index > 92 { //что бы инт не привысить
		return -1
	}

	if len(*sequence)-1 < index {
		*sequence = append(*sequence, make([]int, index-len(*sequence)+1)...)
	}

	if index == 0 {
		(*sequence)[0] = 0
		return 0
	} else if index == 1 {
		(*sequence)[1] = 1
		return 1
	}

	if (*sequence)[index] > 0 {
		return (*sequence)[index]
	}

	(*sequence)[index] = FillSequence(index-1, sequence) + FillSequence(index-2, sequence)
	return (*sequence)[index]
}
