package ascii_test

import (
	"ascii-art/ascii"
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected []string
	}
	tests := []testCase{
		{
			name:     "Single string",
			input:    "../testdata/test.txt",
			expected: []string{" _    _  ", "| |  | | ", "| |__| | ", "|  __  | ", "| |  | | ", "|_|  |_| ", "         "},
		},
	}

	for _, tt := range tests {
		result, err := ascii.ReadFile(tt.input)

		if err != nil {
			t.Errorf("ReadFile failed: %v", err)
		}

		if !reflect.DeepEqual(result, tt.expected) {
			t.Error("Content does not match with expected")
		}
	}

}

func TestAsciiPrinter(t *testing.T) {
	type Parameters struct {
		arr       []string
		fileInput []string
	}
	type testCase struct {
		name     string
		input    Parameters
		expected string
	}
	result, err := ascii.ReadFile("../standard.txt")

	if err != nil {
		t.Errorf("ReadFile failed: %v", err)
	}

	tests := []testCase{
		{
			name:  "test for slice with 1 item",
			input: Parameters{arr: []string{"Hello"}, fileInput: result},
			expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
`,
		},
		{
			name:  "test for empty slice",
			input: Parameters{arr: []string{""}, fileInput: result},
			expected: `
`,
		},
	}

	for _, tt := range tests {
		s := ascii.AsciiPrinter(tt.input.arr, tt.input.fileInput)
		if s != tt.expected {
			t.Errorf("AsciiPrinter function returned\n%q\nexpected -->\n%q", s, tt.expected)
		}
	}
}
