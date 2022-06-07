package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	//var err *customError = nil //print ok
	err = test()
	if err != nil {
		println("error") //error
		return
	}
	println("ok")
}

/*
Функиця test возвращает [nil, *customError], первый элемент - значение,
второй - тип значения. Следовательно [nil, *customError]!=[nil, nil]
*/