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
