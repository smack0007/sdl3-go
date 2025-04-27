package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"sort"
	"strings"
)

var FILES_TO_EXCLUDE = []string{
	"SDL.h",
	"SDL_begin_code.h",
	"SDL_close_code.h",
	"SDL_oldnames.h",
}

var TYPE_MAP = map[string]string{
	"bool":     "bool",
	"char*":    "string",
	"double":   "float64",
	"float":    "float32",
	"float*":   "*float32",
	"Sint32":   "int32",
	"size_t":   "uint64",
	"size_t*":  "*uint64",
	"Uint8":    "uint8",
	"Uint8*":   "*uint8",
	"Uint32":   "uint32",
	"Uint32*":  "*uint32",
	"Uint64":   "uint64",
	"void":     "void",
	"wchar_t*": "string",

	"SDL_PROP_":         "string",
	"SDL_LOG_CATEGORY_": "int",
	"SDL_LOG_PRIORITY_": "LogPriority",
	"SDL_WINDOWPOS_":    "int",
	"SDLK_":             "Keycode",
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
		if slices.Contains(FILES_TO_EXCLUDE, fileName) || strings.HasPrefix(fileName, "SDL_test_") {
			continue
		}
		parseFile(includeDirectory, fileName)
	}
}

func indent(indentLevel int) string {
	return strings.Repeat("\t", indentLevel)
}

func pointerSuffixToPrefix(name string) string {
	for strings.HasSuffix(name, "*") {
		name = "*" + name[0:len(name)-1]
	}

	return name
}

func mapType(name string) string {
	for key, value := range TYPE_MAP {
		if strings.HasPrefix(name, key) {
			return value
		}
	}

	name = pointerSuffixToPrefix(stripPrefixes(name))

	// fmt.Printf("%s\n", name)
	return name
}

func isPointer(name string) bool {
	return strings.HasPrefix(name, "*") || strings.HasSuffix(name, "*")
}

func isString(name string) bool {
	return name == "char*"
}

type ParseState uint32

const (
	ParseStateNone ParseState = iota
	ParseStateEnum
	ParseStateFunc
)

func parseFile(includeDirectory string, fileName string) {
	file, err := os.Open(path.Join(includeDirectory, fileName))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	state := ParseStateNone
	buffer := ""
	defines := []string{}
	enums := []string{}
	funcs := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if state == ParseStateNone {
			if strings.HasPrefix(line, "#define ") && !strings.HasSuffix(line, "_h_") {
				defines = append(defines, line)
			} else if strings.HasPrefix(line, "typedef enum ") {
				state = ParseStateEnum
				buffer += line
			} else if strings.HasPrefix(line, "extern SDL_DECLSPEC ") {
				if strings.Contains(line, ";") {
					funcs = append(funcs, line)
				} else {
					state = ParseStateFunc
					buffer += line
				}
			}
		} else {
			buffer += line

			if strings.Contains(line, ";") {
				switch state {
				case ParseStateEnum:
					enums = append(enums, buffer)
				case ParseStateFunc:
					funcs = append(funcs, buffer)
				}

				buffer = ""
				state = ParseStateNone
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	output := ""

	output = writeDefinesAndEnums(output, defines, enums)
	output = writeFuncs(output, funcs)

	output += "\n"
	outputFileName := strings.ReplaceAll(fileName, "SDL_", "")
	outputFileName = strings.ReplaceAll(outputFileName, ".h", ".nogo")
	err = os.WriteFile(path.Join(".", "tmp", outputFileName), []byte(output), os.FileMode(0644))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func stripPrefixes(input string) string {
	if strings.HasPrefix(input, "SDLK_") {
		return strings.TrimPrefix(input, "SDL")
	}

	return strings.TrimPrefix(input, "SDL_")
}

func getEnumName(input string) string {
	return stripPrefixes(input)
}

func writeDefinesAndEnums(output string, defines []string, enums []string) string {
	for _, enum := range enums {
		enum = minimizeWhitespace(removeComments(strings.ReplaceAll(enum, "typedef enum ", "")))
		parts := trimAllSpace(splitAny(enum, "{,};"))

		if len(parts) < 1 {
			continue
		}

		output += fmt.Sprintf("type %s C.%s\n", getEnumName(parts[0]), parts[0])
	}

	output += "\nconst (\n"

	for _, define := range defines {
		define = minimizeWhitespace(removeComments(strings.ReplaceAll(define, "#define ", "")))
		parts := removeEmptyStrings(strings.Split(define, " "))
		if len(parts) > 1 {
			name := getEnumName(parts[0])

			if strings.Contains(name, "(") && strings.Contains(name, ")") {
				continue
			}

			output += fmt.Sprintf("\t%s %s = C.%s\n", name, mapType(parts[0]), parts[0])
		}
	}

	for _, enum := range enums {
		enum = minimizeWhitespace(removeComments(strings.ReplaceAll(enum, "typedef enum ", "")))
		parts := trimAllSpace(splitAny(enum, "{,};"))

		if len(parts) < 1 {
			continue
		}

		enumName := getEnumName(parts[0])

		for _, enumValue := range parts[1 : len(parts)-1] {
			output += fmt.Sprintf("\t%s %s = C.%s\n", stripPrefixes(enumValue), enumName, parts[0])
		}
	}

	output += ")\n\n"

	return output
}

func writeFuncs(output string, funcs []string) string {
	funcNamesToOutput := make(map[string]string)

	for _, function := range funcs {
		function = minimizeWhitespace(removeComments(strings.ReplaceAll(strings.ReplaceAll(function, "extern SDL_DECLSPEC ", ""), "SDLCALL ", "")))
		function = strings.ReplaceAll(function, "const ", "")
		function = strings.ReplaceAll(strings.ReplaceAll(function, " * ", "* "), " ** ", "** ")
		function = strings.ReplaceAll(function, " *", "* ")

		parts := splitAny(function, " (,);")

		if len(parts) < 2 {
			continue
		}

		returnType := parts[0]
		mappedReturnType := mapType(returnType)
		mapErrorFunc := ""

		funcName := parts[1]

		funcOutput := fmt.Sprintf("func %s(", stripPrefixes(funcName))

		for i := 2; i < len(parts); i += 2 {
			if i+1 >= len(parts) {
				continue
			}
			if i != 2 {
				funcOutput += ", "
			}
			funcOutput += fmt.Sprintf("%s %s", parts[i+1], mapType(parts[i]))
		}

		funcOutput += ") "

		if returnType != "void" {
			returnTypeToOutput := mappedReturnType

			if mappedReturnType == "bool" {
				mapErrorFunc = "BoolToError"
				returnTypeToOutput = "error"
			} else if isPointer(mappedReturnType) {
				mapErrorFunc = "PointerToError"
				returnTypeToOutput = "(" + mappedReturnType + ", error)"
			}

			funcOutput += returnTypeToOutput + " "
		}

		funcOutput += "{\n"

		indentLevel := 1

		// Marshall strings
		for i := 2; i < len(parts); i += 2 {
			if i+1 >= len(parts) {
				continue
			}
			if isString(parts[i]) {
				funcOutput += indent(indentLevel) + "c_" + parts[i+1] + " := C.CString(" + parts[i+1] + ")\n"
				funcOutput += indent(indentLevel) + "defer C.free(unsafe.Pointer(c_" + parts[i+1] + "))\n\n"
			}
		}

		if returnType != "void" {
			funcOutput += indent(indentLevel)

			if mapErrorFunc != "" {
				if mapErrorFunc == "ErrorToBool" {
					funcOutput += "return ErrorToBool(\n"
					indentLevel += 1
					funcOutput += indent(indentLevel)
				} else if mapErrorFunc == "PointerToError" {
					funcOutput += "result := "
				}
			} else {
				funcOutput += "return "
			}

			if isPointer(mappedReturnType) {
				funcOutput += "(" + mappedReturnType + ")(unsafe.Pointer(\n"
			} else {
				funcOutput += mappedReturnType + "(\n"
			}
			indentLevel += 1
		}

		funcOutput += indent(indentLevel) + "C." + funcName + "(\n"
		indentLevel += 1

		for i := 2; i < len(parts); i += 2 {
			if i+1 >= len(parts) {
				continue
			}

			if isString(parts[i]) {
				funcOutput += indent(indentLevel) + "c_" + parts[i+1] + ",\n"
			} else {
				if isPointer(parts[i]) {
					paramType := pointerSuffixToPrefix("C." + parts[i])
					funcOutput += indent(indentLevel) + "(" + paramType + ")(unsafe.Pointer(" + parts[i+1] + ")),\n"
				} else {
					funcOutput += indent(indentLevel) + "C." + parts[i] + "(" + parts[i+1] + "),\n"
				}

			}
		}

		indentLevel -= 1
		funcOutput += indent(indentLevel) + "),\n"

		if returnType != "void" {
			indentLevel -= 1
			funcOutput += indent(indentLevel)

			if mapErrorFunc == "BoolToError" {
				funcOutput += "),\n"
				indentLevel -= 1
				funcOutput += indent(indentLevel) + ")\n"
			} else if mapErrorFunc == "PointerToError" {
				funcOutput += "))\n\n"
				funcOutput += indent(indentLevel) + "return result, PointerToError(result)\n"
			} else {
				funcOutput += ")\n"
			}
		}

		funcOutput += "}\n\n"

		funcNamesToOutput[stripPrefixes(funcName)] = funcOutput
	}

	funcNames := make([]string, 0, len(funcNamesToOutput))
	for key := range funcNamesToOutput {
		funcNames = append(funcNames, key)
	}
	sort.Strings(funcNames)

	for _, funcName := range funcNames {
		output += funcNamesToOutput[funcName]
	}

	return output
}
