package exec_test

import (
	"bytes"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/mdwhatcott/exec"
)

func Test(t *testing.T) {
	_, here, _, _ := runtime.Caller(0)
	output, err := exec.Run("ls -1",
		exec.Options.At(filepath.Dir(here)),
		exec.Options.Out(bytes.NewBufferString("")),
	)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output, filepath.Base(here)) {
		t.Error("missing this file in the listing")
	}
	t.Log("listing:", strings.ReplaceAll(output, "\n", ", "))
}
func TestMust(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("Should have panicked...")
		}
	}()
	_ = exec.MustRun("not-a-chance-in-the-world")
}
func TestJust(t *testing.T) {
	output := exec.JustRun("not-a-chance-in-the-world")
	t.Log(output)
	if !strings.Contains(output, "not found") {
		t.Error("Expected some sort of not-found error from the shell. Output:", output)
	}
}
func TestIn(t *testing.T) {
	output, err := exec.Run("grep 'ello'", exec.Options.In(bytes.NewBufferString("Hello, world!")))
	t.Log(output)
	if output != "Hello, world!" {
		t.Error("missing expected match")
	}
	if err != nil {
		t.Error(err)
	}
}
