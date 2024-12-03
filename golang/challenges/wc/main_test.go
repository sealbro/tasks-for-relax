package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestCase(t *testing.T) {
	content := `The Project
    Release date: May 1, 1994 [eBook #132]
This ebook 

Title: The Art of War`

	filename := "example.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Error("Error creating file:", err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		t.Error("Error writing to file:", err)
	}
	file.Close()

	defer func() {
		os.Remove(filename)
	}()

	replaceStdout := func() (*os.File, *os.File, <-chan string) {
		origin := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		outC := make(chan string)
		go func() {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			outC <- buf.String()
			close(outC)
		}()

		return origin, w, outC
	}

	replaceStdin := func(input string) (*os.File, *os.File) {
		origin := os.Stdin
		r, w, _ := os.Pipe()
		os.Stdin = r

		go func() {
			_, _ = w.Write([]byte(input))
			w.Close()
		}()

		return origin, w
	}

	t.Run("No options and file", func(t *testing.T) {
		originStdout, newStdout, stdoutCh := replaceStdout()

		args := []string{"wc", filename}
		os.Args = args
		main()

		newStdout.Close()
		os.Stdout = originStdout
		out := <-stdoutCh

		expected := "       4      16      89 example.txt\n"
		if out != expected {
			t.Errorf("Expected: %s, got: %s", expected, out)
		}
	})

	t.Run("Options and file", func(t *testing.T) {
		originStdout, newStdout, stdoutCh := replaceStdout()

		args := []string{"wc", "-cw", filename}
		os.Args = args
		main()

		newStdout.Close()
		os.Stdout = originStdout
		out := <-stdoutCh

		expected := "       16      89 example.txt\n"
		if out != expected {
			t.Errorf("Expected: %s, got: %s", expected, out)
		}
	})

	t.Run("Pipe without options and file", func(t *testing.T) {
		originStdout, newStdout, stdoutCh := replaceStdout()
		originStdin, newStdin := replaceStdin(content)

		args := []string{"wc"}
		os.Args = args
		main()

		newStdout.Close()
		os.Stdout = originStdout
		out := <-stdoutCh

		newStdin.Close()
		os.Stdin = originStdin

		expected := "       4      16      89\n"
		if out != expected {
			t.Errorf("Expected: %s, got: %s", expected, out)
		}
	})

	t.Run("Pipe without options and file", func(t *testing.T) {
		originStdout, newStdout, stdoutCh := replaceStdout()
		originStdin, newStdin := replaceStdin(content)

		args := []string{"wc", "-cw"}
		os.Args = args
		main()

		newStdout.Close()
		os.Stdout = originStdout
		out := <-stdoutCh

		newStdin.Close()
		os.Stdin = originStdin

		expected := "       16      89\n"
		if out != expected {
			t.Errorf("Expected: %s, got: %s", expected, out)
		}
	})
}
