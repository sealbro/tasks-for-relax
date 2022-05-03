package goal_parser_interpretation

import "strings"

// https://leetcode.com/problems/goal-parser-interpretation/

func interpret(command string) string {
	builder := strings.Builder{}

	for i := 0; i < len(command); {
		switch command[i] {
		case 'G':
			builder.WriteByte('G')
			i++
		case '(':
			if command[i+1] == ')' {
				builder.WriteByte('o')
				i += 2
			} else if command[i+1] == 'a' {
				builder.WriteByte('a')
				builder.WriteByte('l')
				i += 4
			}
		default:
			i++ // infinity loop escape
		}
	}

	return builder.String()
}
