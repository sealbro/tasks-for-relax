package print_immutable_linked_list_in_reverse

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func toImmutableListNode(input []int) ImmutableListNode {
	if len(input) == 0 {
		return nil
	}

	head := &ImmutableListNodeImp{
		value: input[0],
	}

	current := head
	for i := 1; i < len(input); i++ {
		current.next = &ImmutableListNodeImp{
			value: input[i],
		}
		current = current.next
	}

	return head

}

func TestCase(t *testing.T) {
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

	testCases := []struct {
		input    ImmutableListNode
		expected string
	}{
		{
			input:    toImmutableListNode([]int{1, 2, 3, 4, 5}),
			expected: "5\n4\n3\n2\n1\n",
		},
		{
			input:    toImmutableListNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			expected: "10\n9\n8\n7\n6\n5\n4\n3\n2\n1\n",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase #%d", i+1), func(t *testing.T) {
			originStdout, newStdout, stdoutCh := replaceStdout()

			printLinkedListInReverse(testCase.input)
			newStdout.Close()
			os.Stdout = originStdout
			out := <-stdoutCh

			if out != testCase.expected {
				t.Errorf("Expected: %s, got: %s", testCase.expected, out)
			}
		})
	}
}
