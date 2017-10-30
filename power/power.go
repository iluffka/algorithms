package main

import "fmt"

func main()  {
	fmt.Println(power(5, 120))
}

func power(a, n int) int {
	result := 1 //для хранения результата
	aInDegreeOf2 := a //текущее значение ((a^2)^2...)^2
	for {
		if n <= 0 {
			break
		}
		//добавляем нужную степень двойки к результату,
		//если она есть в разложении n
		if n & 1 == 1 {
			result *= aInDegreeOf2
		}
		aInDegreeOf2 *= aInDegreeOf2
		n = n >> 1 //можно писать n/2
	}
	return result
}

/*
* n%2 == 1 (n&1 == 1) одно и тоже
* n = n >>1 <=> n = n/2
*/
