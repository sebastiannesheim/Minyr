package yr

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const inputFile = "../kjevik-temp-celsius-20220318-20230318.csv"
const outputFile = "../kjevik-temp-fahr-20220318-20230318.csv"

func TestNumberOfLines(t *testing.T) {
	// Test number of lines in input file
	testNumberOfLines(t, inputFile, 16756)

	// Test number of lines in output file
	testNumberOfLines(t, outputFile, 16756)
}

func testNumberOfLines(t *testing.T, filename string, expected int) {
	count, err := GetNumberOfLines(filename)
	if err != nil {
		t.Fatalf("Error counting lines: %v", err)
	}

	if count != expected {
		t.Errorf("Unexpected number of lines in file %s: expected %d, got %d", filename, expected, count)
	}
}

func TestGetNumberOfLines(t *testing.T) {
	// Create a temporary file with some lines
	file, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(file.Name())

	lines := []string{
		"line 1",
		"line 2",
		"line 3",
	}

	for _, line := range lines {
		fmt.Fprintln(file, line)
	}

	// Test that the file has the expected number of lines
	testNumberOfLines(t, file.Name(), len(lines))
}

func TestCelsiusToFahrenheitString(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "6", want: "42.8"},
		{input: "0", want: "32.0"},
	}

	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitString(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}
}

func TestCelsiusToFahrenheitConversion(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
	}

	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitLine(tc.input)
		if tc.want != got {
			t.Errorf("Conversion test failed: expected %s, got: %s", tc.want, got)
		}
	}
}
