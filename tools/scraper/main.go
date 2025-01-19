package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strings"
)

var FILES_TO_EXCLUDE = []string{
	"SDL.h",
	"SDL_begin_code.h",
	"SDL_close_code.h",
	"SDL_oldnames.h",
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the directory to the SDL include files.\n")
		os.Exit(1)
	}

	_ = os.Mkdir(path.Join(".", "tmp"), os.FileMode(0777))

	includeDirectory := path.Join(strings.TrimPrefix(os.Args[1], "-I"), "SDL3")

	entries, err := os.ReadDir(includeDirectory)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: Failed while reading input directory.\n")
		os.Exit(1)
	}

	for _, file := range entries {
		fileName := file.Name()
		if slices.Contains(FILES_TO_EXCLUDE, fileName) {
			continue
		}
		parseFile(includeDirectory, fileName)
	}
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

func minimizeWhitespace(input string) string {
	input = strings.ReplaceAll(input, "\t", " ")
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
		input = substring(input, 0, indexStart) + substring(input, indexEnd+1, len(input)-1)
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

func parseFile(includeDirectory string, fileName string) {
	file, err := os.Open(path.Join(includeDirectory, fileName))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	defines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "#define SDL_") && !strings.HasSuffix(line, "_h_") {
			defines = append(defines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	output := ""

	output += "const (\n"
	for _, define := range defines {
		define = minimizeWhitespace(removeComments(strings.ReplaceAll(define, "#define ", "")))
		parts := removeEmptyStrings(strings.Split(define, " "))
		if len(parts) > 1 {
			name := strings.TrimPrefix(parts[0], "SDL_")
			type_ := "uint32"

			if strings.Contains(name, "(") && strings.Contains(name, ")") {
				continue
			}

			output += fmt.Sprintf("\t%s %s = C.%s\n", name, type_, parts[0])
		}
	}
	output += ")\n"

	output += "\n"
	outputFileName := strings.ReplaceAll(fileName, "SDL_", "")
	outputFileName = strings.ReplaceAll(outputFileName, ".h", ".nogo")
	err = os.WriteFile(path.Join(".", "tmp", outputFileName), []byte(output), os.FileMode(0644))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
