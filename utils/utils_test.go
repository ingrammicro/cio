package utils

import (
	"context"
	"os"
	"reflect"
	"sort"
	"testing"
)

func TestUntar(t *testing.T) {
	tests := map[string]struct {
		source string
		target string
	}{

		"if command success": {
			source: "testdata/sample.zip",
			target: "testdata/uncompressed",
		},
		"if invalid folder": {
			source: "testdata/sample.tar.gz",
			target: "testdata/uncompressed/test_utils",
		},
		"if command failed": {
			source: "testdata/sample.tar.gz",
			target: "testdata/uncompressed",
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			err := Untar(context.Background(), test.source, test.target)
			if err != nil && test.source == "testdata/sample.zip" {
				t.Errorf("Unexpected error: %v\n", err)
			}
			os.Remove("testdata/uncompressed")
		})
	}
}

func TestFileExists(t *testing.T) {
	tests := map[string]struct {
		name     string
		expected bool
	}{
		"if folder exists": {
			name:     "testdata",
			expected: true,
		},
		"if folder does not exists": {
			name:     "test_utils",
			expected: false,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			bExist := FileExists(test.name)
			if bExist != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", bExist, test.expected)
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	tests := map[string]struct {
		len int
	}{
		"if random string is generated": {
			len: 20,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			result := RandomString(test.len)
			if result == "" {
				t.Error("No empty random string generated\n")
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := map[string]struct {
		elements []string
		expected []string
	}{
		"if remove duplicate elements in a slice": {
			elements: []string{"a", "a", "b", "c", "d", "e"},
			expected: []string{"a", "b", "c", "d", "e"},
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			result := RemoveDuplicates(test.elements)
			t.Logf("%v\n", result)
			sort.Strings(result)
			sort.Strings(test.expected)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", result, test.expected)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := map[string]struct {
		chain     []string
		character string
		expected  bool
	}{
		"if the slice contains the character": {
			chain:     []string{"a", "b", "c", "d", "e"},
			character: "a",
			expected:  true,
		},
		"if the slice does not contains the character": {
			chain:     []string{"a", "b", "c", "d", "e"},
			character: "z",
			expected:  false,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			bOk := Contains(test.chain, test.character)
			if bOk != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", bOk, test.expected)
			}
		})
	}
}
