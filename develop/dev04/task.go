package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

//sortWord - returns the word sorted alphabetically
func sortWord(word string) string {
	arr := []rune(word)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return string(arr)
} 

//sortStringSlice - returns words sorted alphabetically
func sortStringSlice(strs []string) []string {
	sort.Slice(strs, func(i, j int) bool { return strs[i] < strs[j] })
	return strs
}

//removeDuplicates - removes all duplicates in an array 
func removeDuplicates(strs []string) []string {
	if len(strs) == 0 {
		return nil
	}
	var res []string
	for i := 1; i < len(strs); i++ {
		if strs[i] != strs[i - 1] {
			res = append(res, strs[i - 1])
		}
	}
	res = append(res, strs[len(strs) - 1])
	return res
}

//SearchAnagrams - searches for all anagrams in an array
func SearchAnagrams(vocabulary []string) map[string][]string {

	mapDuplicate := make(map[string][]int)
	mapResult := make(map[string][]string)

	//sort words by letters and memorize their indexes
	for i, word := range vocabulary {
		sWord := sortWord(word)
		mapDuplicate[sWord] = append(mapDuplicate[sWord], i)
	}

	for mapDupK := range mapDuplicate {
		if len(mapDuplicate[mapDupK]) > 1 {
			for _, numberWord := range mapDuplicate[mapDupK] {
				mapResult[vocabulary[mapDuplicate[mapDupK][0]]] = append(mapResult[vocabulary[mapDuplicate[mapDupK][0]]], vocabulary[numberWord])
			}
		}
	}


	//sort and remove duplicates
	for mapR := range mapResult {
		mapResult[mapR] = sortStringSlice(mapResult[mapR])
		mapResult[mapR] = removeDuplicates(mapResult[mapR])
	}

	return mapResult
}

func main() {

	//open file
	file, err := os.Open("anagrams.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var vocabulary []string

	//read file line by line
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		vocabulary = append(vocabulary, strings.ToLower(fileScanner.Text()))
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(vocabulary)

	anagrams := SearchAnagrams(vocabulary)

	fmt.Println(anagrams)

}
