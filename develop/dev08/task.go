package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
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
//type forkS struct{} 
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

func (shell *echoS) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	var output bytes.Buffer
	var err error

	if len(args) > 0 {
		_, err = output.WriteString(args[0] + "\n")
	} else {
		_, err = output.WriteString("\n")
	}

	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (shell *killS) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	if len(args) > 0 {
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, err
		}
		syscall.Kill(pid, syscall.SIGKILL)
	}
	return nil, nil
}

type proc struct {
	name string
	pid int
}

func (shell *psS) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	var output bytes.Buffer

	ps, err := filepath.Glob("/proc/*/exe")
	if err != nil {
		return nil, err
	}

	var procs []proc
	for _, file := range ps {
		link, _ := os.Readlink(file)
		if len(link) > 0 {
			name := filepath.Base(link)
			pid, err := strconv.Atoi(strings.Split(file, "/")[2])

			if err == nil {
				procs = append(procs, proc{name, pid})
			}
		}
	}

	sort.Slice(procs, func(i, j int) bool { return procs[i].pid < procs[j].pid })

	output.WriteString("PID Name\n")
	for _, val := range procs {
		output.WriteString(fmt.Sprintf("%d %s\n", val.pid, val.name))
	}
	return &output, nil
}


func (shell *execS) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	com := exec.Command(args[0], args[1:]...)
	var output bytes.Buffer

	if input != nil {
		com.Stdin = bytes.NewReader((*input).Bytes())
	}

	com.Stdout = &output
	com.Stderr = os.Stderr
	if err := com.Run(); err != nil {
		return nil, err
	}

	return &output, nil
}

func (shell *ncS) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	if len(args) < 3 {
		return nil, errors.New("too few args (type, ip, port)")
	}

	conn, err := net.Dial(args[0], args[1] + ":" + args[2])
	if err != nil {
		return nil, err
	}

	var str string
	fmt.Printf(">")
	fmt.Scan(&str)

	_, err = conn.Write([]byte(str)) 

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (shell *quitS) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	if !pipes {
		os.Exit(0)
	}
	return nil, nil
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
	case "kill":
		sh = &killS{}
	case "ps":
		sh = &psS{}
	case "exec":
		sh = &execS{}
	case "nc":
		sh = &ncS{}
	case "quit":
		sh = &quitS{}
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
