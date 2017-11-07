package cmd

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"testing"
)

type TestData struct {
	Name string `json:"NameTag"`
	Data string
}

func TestPrintResponse(t *testing.T) {

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	data := TestData{"Test", "Test Data"}
	PrintResponse(data)

	expected := `+--------+----------+
|NameTag |          |
+--------+----------+
|Test    |Test Data |
+--------+----------+
`

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if out != expected {
		t.Error("Print Response did not match expected output\n", expected, "\nactual\n", out)
	}

}

func TestHandleErrorNil(t *testing.T) {
	HandleError(nil)
}

func TestHandleError(t *testing.T) {

	if os.Getenv("BE_CRASHER") == "1" {
		HandleError(errors.New("test error"))
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestHandleError")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)

}
