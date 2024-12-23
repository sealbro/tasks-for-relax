package print_foobar_alternately

import (
	"bytes"
	"fmt"
	"golang/assert"
	"io"
	"math/rand/v2"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		input int
	}{
		{input: 1},
		{input: 5},
		{input: 10},
	}

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

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			originStdout, newStdout, stdoutCh := replaceStdout()

			fb := NewFooBar(tc.input)

			wg := sync.WaitGroup{}
			wg.Add(tc.input * 2)

			go func() {
				fb.Foo(func() {
					defer wg.Done()
					time.Sleep(time.Millisecond * time.Duration(rand.IntN(200)))
					fmt.Print("foo")
				})
			}()

			go func() {
				fb.Bar(func() {
					defer wg.Done()
					time.Sleep(time.Millisecond * time.Duration(rand.IntN(200)))
					fmt.Print("bar")
				})
			}()

			wg.Wait()

			newStdout.Close()
			os.Stdout = originStdout
			actual := <-stdoutCh

			sb := strings.Builder{}
			for i := 0; i < tc.input; i++ {
				sb.WriteString("foobar")
			}
			expected := sb.String()

			assert.Equal(t, expected, actual)
		})
	}
}
