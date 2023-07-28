package exec

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
)

var Options opt

type option func(*exec.Cmd)

type opt struct{}

func (opt) At(dir string) option {
	return func(command *exec.Cmd) {
		command.Dir = dir
	}
}

func (opt) Out(writers ...io.Writer) option {
	return func(command *exec.Cmd) {
		command.Stdout = io.MultiWriter(append(writers, command.Stdout)...)
		command.Stderr = command.Stdout
	}
}

func Run(program string, options ...option) (output string, err error) {
	buffer := new(bytes.Buffer)
	command := exec.Command("bash", "-c", program)
	command.Stdout = buffer
	command.Stderr = buffer
	for _, option := range options {
		option(command)
	}
	err = command.Run()
	return strings.TrimSpace(buffer.String()), err
}
