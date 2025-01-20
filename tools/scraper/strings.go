package main

import (
	"strings"
)

func minimizeWhitespace(input string) string {
	input = strings.ReplaceAll(input, "\t", " ")
	input = strings.ReplaceAll(input, ", ", ",")
	input = strings.ReplaceAll(input, "  ", " ")
	return input
}

func removeComments(input string) string {
	index := strings.Index(input, "//")
	if index != -1 {
		input = input[0:index]
	}

	indexStart := strings.Index(input, "/*")
	indexEnd := strings.Index(input, "*/")
	if indexStart != -1 && indexEnd != -1 {
		result := ""
		output := true
		for i := 0; i < len(input); i += 1 {
			if output {
				if i < len(input)-1 && input[i] == '/' && input[i+1] == '*' {
					output = false
				} else {
					result += string(input[i])
				}
			} else {
				if i < len(input)-1 && input[i] == '*' && input[i+1] == '/' {
					output = true
					i += 1
				}
			}
		}

		input = result
	}

	return input
}

func removeEmptyStrings(input []string) []string {
	result := []string{}
	for _, value := range input {
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}

func splitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func substring(input string, start, end int) string {
	counter, startIndex := 0, 0
	for i := range input {
		if counter == start {
			startIndex = i
		}
		if counter == end {
			return input[startIndex:i]
		}
		counter += 1
	}
	return input[startIndex:]
}

func trimAllSpace(input []string) []string {
	result := make([]string, len(input))

	for index, value := range input {
		result[index] = strings.TrimSpace(value)
	}

	return result
}
