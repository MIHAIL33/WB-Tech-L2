package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s, cap(s)) // [3 2 3]
}

func modifySlice(i []string) {
	i[0] = "3"
	fmt.Println("Modify cap:", cap(i))
	i = append(i, "4")
	fmt.Println("Modify cap:", cap(i))
	i[1] = "5"
	i = append(i, "6")
}

/*
Cлайс s имеет capacity равное количеству элементов при инициализации (capacity = 3).
Т.к слайс хранит внутри себя ссылку на массив, при превышении capacity массива
произойдет аллокация, и вернется ссылка на новый массив. 
*/