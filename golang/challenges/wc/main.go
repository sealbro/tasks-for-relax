package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type Options int

const (
	ByteCount = iota
	LineCount
	WordCount
	CharCount
)

// https://codingchallenges.fyi/challenges/challenge-wc
func main() {
	optionsArgs, fileNameArgs := argumentsParser(os.Args)
	options, err := parseOptions(optionsArgs)
	if err != nil {
		fmt.Println("Error parsing options:", err)
		return
	}

	var file *os.File
	if len(fileNameArgs) == 0 {
		file, err = openPipe()
		if err != nil {
			fmt.Println("Error opening pipe:", err)
			return
		}
		defer fmt.Println("")
	} else {
		var bytes uint64
		file, bytes, err = openFile(fileNameArgs)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		options[ByteCount] = bytes
	}
	defer file.Close()

	result, err := parseFile(file, options)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// Output
	intend := strings.Repeat(" ", 6)
	fmt.Print(" ")
	if _, ok := result[LineCount]; ok {
		fmt.Print(intend)
		fmt.Print(result[LineCount])
	}
	if _, ok := result[WordCount]; ok {
		fmt.Print(intend)
		fmt.Print(result[WordCount])
	}
	if _, ok := result[ByteCount]; ok {
		fmt.Print(intend)
		fmt.Print(result[ByteCount])
	}
	if fileNameArgs != "" {
		fmt.Print(" ")
		fmt.Println(fileNameArgs)
	}
}

func openFile(fileName string) (*os.File, uint64, error) {
	stat, err := os.Stat(fileName)
	if err != nil {
		return nil, 0, err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, 0, err
	}

	return file, uint64(stat.Size()), nil
}

func openPipe() (*os.File, error) {
	return os.Stdin, nil
}

func argumentsParser(args []string) (options, fileName string) {
	if len(args) == 1 {
		return "", ""
	}

	if len(args) == 2 {
		if args[1][0] == '-' {
			return args[1], ""
		} else {
			return "", args[1]
		}
	}

	return args[1], args[2]
}

func parseFile(reader io.Reader, options map[Options]uint64) (map[Options]uint64, error) {
	_, hasWordCount := options[WordCount]
	_, hasCharCount := options[CharCount]
	_, hasLineCount := options[LineCount]

	// Stdin case
	bytesValue, hasBytesCount := options[ByteCount]
	if bytesValue > 0 {
		hasBytesCount = false
	}

	if hasWordCount || hasCharCount || hasLineCount {
		buffer := make([]byte, 1024)
		var g uint64 = 0
		var lastSpace uint64 = 0
		for {
			n, err := reader.Read(buffer)
			if err != nil {
				if errors.Is(err, io.EOF) {
					if lastSpace+1 != g {
						options[WordCount] += 1
					}
					break
				}
				return nil, err
			}
			// If we read 0 bytes, we are done
			if n == 0 {
				break
			}

			if hasWordCount || hasBytesCount {
				for _, b := range buffer[:n] {
					if hasWordCount {
						r := rune(b)
						if unicode.IsSpace(r) {
							// If the last character was a space, we don't count it
							if lastSpace+1 != g {
								options[WordCount] += 1
							}
							lastSpace = g
						}
						if hasLineCount && b == '\n' {
							options[LineCount] += 1
						}
					}
					g++
				}
			}
			if hasCharCount {
				options[CharCount] += uint64(len(buffer))
			}
		}

		options[ByteCount] = g

	}

	return options, nil
}

func parseOptions(args string) (map[Options]uint64, error) {
	if len(args) == 0 || args[0] != '-' {
		return map[Options]uint64{
			ByteCount: 0,
			LineCount: 0,
			WordCount: 0,
			CharCount: 0,
		}, nil
	}

	var options = make(map[Options]uint64)
	for _, opt := range args[1:] {
		switch opt {
		case 'c':
			options[ByteCount] = 0
		case 'l':
			options[LineCount] = 0
		case 'w':
			options[WordCount] = 0
		case 'm':
			options[CharCount] = 0
		default:
			return nil, fmt.Errorf("invalid option provided")
		}
	}

	if len(options) == 0 {
		return nil, fmt.Errorf("no valid options provided")
	}

	return options, nil
}
