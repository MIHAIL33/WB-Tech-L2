package main

import (
	"errors"
	"fmt"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

//Unpacking - return unpacking string with support escape option
func Unpacking(str string) (*string, error) {
	runes := []rune(str)
 	var result []rune

	var count int
	var slash bool

	for i, runee := range runes {
		if unicode.IsNumber(runee) && i == 0 {
			return nil, errors.New("bad string")
		}

		if unicode.IsNumber(runee) && unicode.IsNumber(runes[i - 1]) && runes[i - 2] != '\\' {
			return nil, errors.New("bad string")
		}

		if runee == '\\' && !slash {
			slash = true
			continue
		}

		if slash {
			result = append(result, runee)
			slash = false
			continue
		}

		if unicode.IsNumber(runee) {
			count = int(runee - '0')
			if count == 0 {
				continue
			}
			for j := 0; j < count - 1; j++ {
				result = append(result, runes[i - 1])
			} 
			continue
		}

		result = append(result, runee)
	}

	res := string(result)
	return &res, nil
}

func main() {
	res1, _ := Unpacking("a4bc2d5e") // aaaabccddddde
	fmt.Println(*res1)

	res2, _ := Unpacking("abcd") // abcd
	fmt.Println(*res2)

	_, err := Unpacking("54") // error
	fmt.Println(err)

	res4, _ := Unpacking("") // ""
	fmt.Println(*res4)

	//with "/"

	res5, _ := Unpacking("qwe\\4\\5") // qwe45
	fmt.Println(*res5)

	res6, _ := Unpacking("qwe\\45") // qwe44444
	fmt.Println(*res6)

	res7, _ := Unpacking("qwe\\\\5") // qwe\\\\\
	fmt.Println(*res7)
}
