package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err) //nil
	fmt.Println(err == nil) //false

	fmt.Println(err == (*os.PathError)(nil)) //true
}

/*
Значение интерфейса равно nil только в том случае, если и его значение,
и динамический тип равны nil. В приведенном выше примере Foo()возвращается,
[nil, *os.PathError] и мы сравниваем его с [nil, nil].
*/