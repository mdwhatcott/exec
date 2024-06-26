package exec

import (
	"bytes"
	"io"
	"os"
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
func (opt) In(reader io.Reader) option {
	return func(command *exec.Cmd) {
		command.Stdin = reader
	}
}
func (opt) Out(writers ...io.Writer) option {
	return func(command *exec.Cmd) {
		command.Stdout = io.MultiWriter(append(writers, command.Stdout)...)
		command.Stderr = command.Stdout
	}
}
func (opt) Env(key, value string) option {
	return func(command *exec.Cmd) {
		command.Env = append(command.Env, key+"="+value)
	}
}

func Run(program string, options ...option) (output string, err error) {
	buffer := new(bytes.Buffer)
	command := exec.Command("bash", "-c", program)
	command.Stdout = buffer
	command.Stderr = buffer
	command.Env = os.Environ()
	for _, option := range options {
		option(command)
	}
	err = command.Run()
	return strings.TrimSpace(buffer.String()), err
}
func MustRun(program string, options ...option) string {
	output, err := Run(program, options...)
	if err != nil {
		panic(err)
	}
	return output
}
func JustRun(program string, options ...option) string {
	output, _ := Run(program, options...)
	return output
}
