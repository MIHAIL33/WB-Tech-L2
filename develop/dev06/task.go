package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func delimiter(strs []string, field int, delim string, sep bool) ([]string, error) {
	var res []string
	var temp []string

	for _, val := range strs {
		temp = strings.Split(val, delim)
		if sep && len(temp) < 2 {
			continue
		}
		if (len(temp) >= field) {
			res = append(res, temp[field - 1])
		}
	}

	return res, nil
}

func print(strs []string) {
	for _, val := range strs {
		fmt.Println(val) 
	}
}

func main() {
	fFlag := flag.Int("f", 1, "fields")
	dFlag := flag.String("d", "\t", "delimiter")
	sFlag := flag.Bool("s", false,  "separated")
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
		strs = append(strs, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	result, err := delimiter(strs, *fFlag, *dFlag, *sFlag)
	if err != nil {
		log.Fatal(err)
	}

	print(result)
}
