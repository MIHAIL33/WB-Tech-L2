package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type sortN struct {
	number int 
	str string
}

func SortNumber(strs []string, isOrder bool) ([]string, error) {
	reg:= regexp.MustCompile(`(^\d+)`)

	var strAccor []sortN
	var numberStr []string
	var badStr []string
	var num int
	var err error

	for _, str := range strs {
		numberStr = reg.FindStringSubmatch(str)
		if numberStr == nil {
			badStr = append(badStr, str)
			continue
		} else {
			num, err = strconv.Atoi(numberStr[0])
			if err != nil {
				return nil, err
			}
		}
		strAccor = append(strAccor, sortN{ number: num, str: str })
	}

	if isOrder {
		sort.SliceStable(strAccor, func(i, j int) bool { return strAccor[i].number < strAccor[j].number })
	} else {
		sort.SliceStable(strAccor, func(i, j int) bool { return strAccor[i].number > strAccor[j].number })
	}
	
	
	var result []string

	for _, sAccor := range strAccor {
		result = append(result, sAccor.str)
	}

	result = append(result, badStr...)

	return result, nil
}

//removeDuplicates - removes all duplicates in an array 
func RemoveDuplicates(strs []string) []string {
	mapStr := make(map[string][]int)
	for i, str := range strs {
		mapStr[str] = append(mapStr[str], i)
	}

	var idxs []int
	for _, val := range mapStr {
		idxs = append(idxs, val[0])
	}

	sort.SliceStable(idxs, func(i, j int) bool { return idxs[i] < idxs[j] })
	var res []string

	for _, val := range idxs {
		res = append(res, strs[val])
	}

	return res
}

func PrintString(strs []string) {
	for _, val := range strs {
		fmt.Println(val)
	}
}

func main() {

	nFlag := flag.Bool("n", false, "sort by numeric value")
	rFlag := flag.Bool("r", false, "sort in reverse order")
	uFlag := flag.Bool("u", false, "do not output duplicate lines")
	flag.Parse()

	fileName := flag.Arg(0)

	//open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var strs []string

	//read file line by line
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		strs = append(strs, strings.ToLower(fileScanner.Text()))
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(strs)

	//strs, _ = SortNumber(strs, false)

	if *uFlag {
		strs = RemoveDuplicates(strs)
	}

	if *nFlag {
		strs, err = SortNumber(strs, !*rFlag)
		if err != nil {
			log.Fatal(err)
		}
	}

	PrintString(strs)
}
