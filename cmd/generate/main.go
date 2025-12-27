package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

const (
	FILE = `package %s

func %s(s string) int {
	out := 0

	return out
}

`
	TEST_FILE = `package %s

import (
	"testing"
)

type TestCase struct {
	Input  string
	Output int
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{Input: "", Output: 0},
	}
	for _, tc := range testCases {
		result := %s(tc.Input)
		if result != tc.Output {
			t.Logf("\nInput: %%+v\nExpected: %%+v\nActual: %%+v", tc.Input, tc.Output, result)
			t.Fail()
		}
	}
}

`
)

func main() {
	if len(os.Args) != 2 {
		slog.Error("Usage: go run generate/main.go <name of problem>")
		return
	}
	name := os.Args[1]
	splitted := split(name)
	functionName := getFunctionName(splitted)
	packageName := getPackageName(splitted)

	fileContents := fmt.Sprintf(FILE, packageName, functionName)
	testFileContents := fmt.Sprintf(TEST_FILE, packageName, functionName)

	err := os.Mkdir(fmt.Sprintf("internal/%s", packageName), 0755)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(fmt.Sprintf("internal/%s/%s.go", packageName, packageName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n, err := file.WriteString(fileContents)
	if err != nil || n != len(fileContents) {
		panic(err)
	}

	file, err = os.Create(fmt.Sprintf("internal/%s/%s_test.go", packageName, packageName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n, err = file.WriteString(testFileContents)
	if err != nil || n != len(testFileContents) {
		panic(err)
	}
}

func split(s string) []string {
	hasSpace := strings.Contains(s, " ")
	hasDash := strings.Contains(s, "-")
	if !hasSpace && !hasDash {
	}
	if hasSpace {
		return strings.Split(s, " ")
	} else if hasDash {
		return strings.Split(s, "-")
	}
	panic(fmt.Sprintf("no space or dash in arg: %s", s))
}

func getPackageName(s []string) string {
	allLower := make([]string, len(s))
	for i := range s {
		allLower[i] = strings.ToLower(s[i])
	}
	return strings.Join(allLower, "_")
}

func getFunctionName(s []string) string {
	allTitle := make([]string, len(s))
	for i := range s {
		allTitle[i] = titleCase(s[i])
	}
	return strings.Join(allTitle, "")
}

func titleCase(s string) string {
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

