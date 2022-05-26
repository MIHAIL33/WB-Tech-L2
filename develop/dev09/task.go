package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func CheckUrl(url string) bool {
	if url != "" {
		runes := []rune(url)
		return runes[0] == '/'
	}
	return false
}


func CreateFile(path string, body io.Reader) error {
	parts := strings.Split(path, "/")
	var tempPath string
	var lenPath int = len(parts)
	for i, val := range parts {

		if i == 0 {
			tempPath = val
		} else {
			tempPath = tempPath + "/" + val
		}

		if i != lenPath - 1 {
			if _, err := os.Stat(tempPath); os.IsNotExist(err) {
				err := os.Mkdir(tempPath, 0750)
				if err != nil {
					return err
				}
			}
		} else {
			file, err := os.Create(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(file, body)
			if err != nil {
				return err
			}
		}

	}

	return nil
}


func DownloadSite(url string, path string) error {

	fmt.Println("Downloading ", url, " to ", path)

	urlWithoutDomain, err := ParseUrl(url, false)
	if err != nil {
		return err
	}

	domain, err := ParseUrl(url, true)
	if err != nil {
		return err
	} 

	mapPath := make(map[string]bool)
	mapPath[urlWithoutDomain] = false

	_, err = DownloadPages(domain, path, urlWithoutDomain, mapPath)
	if err != nil {
		return err
	}

	return nil
}


func DownloadPages(url, path, urlWithoutDomain string, mapPath map[string]bool) (map[string]bool, error) {

	mapPath[urlWithoutDomain] = true

	client := http.Client{}
	req, _ := http.NewRequest("GET", url + urlWithoutDomain, nil)
	req.Header = http.Header{
		"User-Agent": {"a"}, //fix response 410 from Apache
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	if urlWithoutDomain == "/" {
		err = CreateFile(path + urlWithoutDomain + "index.html", bytes.NewBuffer(body))
	} else {
		err = CreateFile(path + urlWithoutDomain + ".html", bytes.NewBuffer(body))
	}
	if err != nil {
		return nil, err
	}

	mapPath = GetLinks(bytes.NewBuffer(body), mapPath)
	
	for k := range mapPath {
		if !mapPath[k] {
			mapPath, err = DownloadPages(url, path, k, mapPath)
			if err != nil {
				return nil, err
			}
		}	
	}

	return mapPath, nil
}


func GetLinks(body io.Reader, mapPath map[string]bool) map[string]bool {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return mapPath
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						if CheckUrl(attr.Val) && !mapPath[attr.Val] {
							mapPath[attr.Val] = false
						}
					}
				}
			} 
		}
	}
}

func ParseUrl(url string, domain bool) (string, error) {
	if url == "" { return "", errors.New("bad url") }
	var localUrl string = ""
	splitFn := func(c rune) bool {
		return c == '/'
	}
	parts := strings.FieldsFunc(url, splitFn)
	if domain {
		localUrl = parts[0] + "//" + parts[1]
		return localUrl, nil
	} else {
		if len(parts) > 2 {
			for i := 2; i < len(parts); i++ {
				localUrl = localUrl + "/" + parts[i]
			}
			return localUrl, nil
		} else {
			return "/", nil
		}
	}
}

func DownloadFile(url string, path string) error {
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"User-Agent": {"a"}, //fix response 410 from Apache
	}

	response, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)

	return err
}

func main() {

	mFlag := flag.Bool("m", false, "download site (mirror)")

	flag.Parse()

	url := flag.Arg(0)

	fmt.Println(url)

	if url == "" {
		fmt.Fprintf(os.Stderr, "Error: empty URL!\n")
		return
	}

	if *mFlag {
		pathLocal, err := ParseUrl(url, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: bad URL!\n")
		}
		pathLocal = path.Base(pathLocal)
		fmt.Println("Checking if " + pathLocal + " exists ...")
		if _, err := os.Stat(pathLocal); os.IsNotExist(err) {
			err := DownloadSite(url, pathLocal)
			if err != nil {
				panic(err)
			}
			fmt.Println(pathLocal + " saved!")
		} else {
			fmt.Println(pathLocal + " already exist!")
		}
	} else {
		pathLocal := path.Base(url)
		fmt.Println("Checking if " + pathLocal + " exists ...")
		if _, err := os.Stat(pathLocal); os.IsNotExist(err) {
			err := DownloadFile(url, pathLocal)
			if err != nil {
				panic(err)
			}
			fmt.Println(pathLocal + " saved!")
		} else {
			fmt.Println(pathLocal + " already exist!")
		}
	}
		
}
