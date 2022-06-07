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

//SortStrs - global type for sorting
type SortStrs struct {
	index int
	str string
}

type sortN struct {
	number int 
	str SortStrs
}

//SortNumber - sorts in ascending or descending order
func SortNumber(strs []SortStrs, isOrder bool) ([]SortStrs, error) {
	reg:= regexp.MustCompile(`(^\d+)`)

	var strAccor []sortN
	var numberStr []string
	var badStr []SortStrs
	var num int
	var err error

	for _, str := range strs {
		numberStr = reg.FindStringSubmatch(str.str)
		if numberStr == nil {
			badStr = append(badStr, SortStrs{str.index, str.str})
			continue
		} else {
			num, err = strconv.Atoi(numberStr[0])
			if err != nil {
				return nil, err
			}
		}
		strAccor = append(strAccor, sortN{ number: num, str: SortStrs{str.index, str.str} })
	}

	if isOrder {
		sort.SliceStable(strAccor, func(i, j int) bool { return strAccor[i].number < strAccor[j].number })
	} else {
		sort.SliceStable(strAccor, func(i, j int) bool { return strAccor[i].number > strAccor[j].number })
	}
	
	
	var result []SortStrs

	for _, sAccor := range strAccor {
		result = append(result, SortStrs{sAccor.str.index, sAccor.str.str})
	}

	result = append(result, badStr...)

	return result, nil
}

//RemoveDuplicates - removes all duplicates in an array 
func RemoveDuplicates(strs []SortStrs) []SortStrs {
	mapStr := make(map[string][]int)
	for i, str := range strs {
		mapStr[str.str] = append(mapStr[str.str], i)
	}

	var idxs []int
	for _, val := range mapStr {
		idxs = append(idxs, val[0])
	}

	sort.SliceStable(idxs, func(i, j int) bool { return idxs[i] < idxs[j] })
	var res []SortStrs

	for _, val := range idxs {
		res = append(res, SortStrs{strs[val].index, strs[val].str})
	}

	return res
}

//PrintString - print all strings
func PrintString(strs []SortStrs) {
	for _, val := range strs {
		fmt.Println(val.str)
	}
}

//split - get only one column
func split(strs []SortStrs, col int) ([]SortStrs, []SortStrs) {
	var res []SortStrs
	var tempSplit []string
	var bad []SortStrs
	for _, val := range strs {
		tempSplit = strings.Split(val.str, " ")
		if len(tempSplit) > col {
			res = append(res, SortStrs{val.index, tempSplit[col]})
		} else {
			bad = append(bad, SortStrs{val.index, val.str})
		}
	}
	return res, bad
}

//sortMatch - match after sorting by column
func sortMatch(strs []SortStrs, newStr []SortStrs, bad []SortStrs) []SortStrs {
	var res []SortStrs
	for _, val := range newStr {
		res = append(res, strs[val.index])
	}

	res = append(res, bad...)
	return res 
}

func main() {

	nFlag := flag.Bool("n", false, "sort by numeric value")
	rFlag := flag.Bool("r", false, "sort in reverse order")
	uFlag := flag.Bool("u", false,  "do not output duplicate lines")
	colFlag := flag.Int("k", 0, "sortable column")

	flag.Parse()

	fileName := flag.Arg(0)

	//open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var strs []SortStrs

	//read file line by line
	fileScanner := bufio.NewScanner(file)
	indx := 0
	for fileScanner.Scan() {
		strs = append(strs, SortStrs{indx, strings.ToLower(fileScanner.Text())})
		indx++
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	if *colFlag > 0 { //sorting by column
		newStr, bad := split(strs, *colFlag)
		if *nFlag {
			newStr, err = SortNumber(newStr, !*rFlag)
			if err != nil {
				log.Fatal(err)
			}
		}
		strs = sortMatch(strs, newStr, bad)
	} else {
		if *nFlag {
			strs, err = SortNumber(strs, !*rFlag)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if *uFlag { //remove duplicates
		strs = RemoveDuplicates(strs)
	}

	PrintString(strs)
}
