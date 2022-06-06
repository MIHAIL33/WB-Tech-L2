package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type shell interface {
	run([]string, *bytes.Buffer, bool) (*bytes.Buffer, error)
}

type cdS struct{}
type pwdS struct{}
type echoS struct{}
type killS struct{}
type psS struct{}
type forkS struct{}
type execS struct{}
type ncS struct{}
type quitS struct{}


func (shell *cdS) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	if !pipes {
		if len(args) > 0 {
			return nil, os.Chdir(args[0])
		}
	}
	return nil, nil
}


func (shell *pwdS) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	var output bytes.Buffer

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	output.WriteString(dir + "\n")
	return &output, nil
}

func (shell *pwdS) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {

}

func runCommand(command string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	
	argsSplit := strings.Split(command, " ")
	var args []string

	for _, val := range argsSplit {
		if val != "" {
			args = append(args, val)
		}
	}

	var sh shell
	switch args[0] {
	case "cd":
		sh = &cdS{}
	case "pwd":
		sh = &pwdS{}
	case "echo":
		sh = &echoS{}
	}

	if sh != nil {
		return sh.run(args[1:], input, pipes)
	}

	return nil, nil
}

func execCommands(commands string) error {
	coms := strings.Split(commands, "|")
	pipes := len(coms) > 1

	var input *bytes.Buffer
	var err error

	for _, val := range coms {
		input, err = runCommand(val, input, pipes)
		if err != nil {
			return err
		}
	}

	if input == nil || len(input.String()) == 0 {
		fmt.Fprint(os.Stdout, "")
	} else {
		fmt.Fprint(os.Stdout, input.String())
	}

	return err
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s$ ", dir)
		commands, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		commands = strings.TrimSuffix(commands, "\n")

		err = execCommands(commands)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}

}
