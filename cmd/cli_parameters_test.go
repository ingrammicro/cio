package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestFlagConvertParamsJSON(t *testing.T) {
	tests := map[string]struct {
		name     string
		data     interface{}
		expected interface{}
	}{
		"if flag is undefined": {
			name:     "",
			data:     nil,
			expected: new(map[string]interface{}),
		},
		"if flag is defined and with a valid value": {
			name:     "valid param",
			data:     map[string]int{"a": 1, "b": 2, "c": 3},
			expected: new(map[string]interface{}),
		},
		"if flag is defined but with an invalid value": {
			name:     "invalid param",
			data:     "a",
			expected: "flag invalid param",
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.data != nil {
				data := test.data
				if reflect.TypeOf(test.data).Name() != "string" {
					data, _ = json.Marshal(test.data)
				}
				viper.Set(test.name, data)
			}

			i, err := FlagConvertParamsJSON(test.name)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && reflect.TypeOf(i) != reflect.TypeOf(test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", i, test.expected)
			}
		})
	}
}

func TestConvertFlagParamsJsonFromFileOrStdin(t *testing.T) {
	tests := map[string]struct {
		param    string
		expected interface{}
	}{
		"if defined input with valid data": {
			param:    "testdata/input_sample.txt",
			expected: map[string]interface{}{"a": "1", "b": "2", "c": "3"},
		},
		"if defined invalid input": {
			param:    "testdata/input_sample_whatever.txt",
			expected: "cannot open",
		},
		"if defined input with no data": {
			param:    "testdata/input_sample_empty.txt",
			expected: "cannot read JSON params from",
		},
		"if defined input as STDIN": {
			param:    "-",
			expected: map[string]interface{}{"a": "1", "b": "2", "c": "3"},
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.param == "-" {
				content := []byte("{\"a\": \"1\", \"b\": \"2\", \"c\": \"3\"}")
				tmpfile, err := ioutil.TempFile("", "example")
				if err != nil {
					t.Errorf("Unexpected error: %v\n", err)
				}
				defer os.Remove(tmpfile.Name()) // clean up

				if _, err := tmpfile.Write(content); err != nil {
					t.Errorf("Unexpected error: %v\n", err)
				}
				if _, err := tmpfile.Seek(0, 0); err != nil {
					t.Errorf("Unexpected error: %v\n", err)
				}
				oldStdin := os.Stdin
				defer func() { os.Stdin = oldStdin }() // Restore original Stdin
				os.Stdin = tmpfile
			}

			i, err := ConvertFlagParamsJsonFromFileOrStdin(test.param)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && reflect.TypeOf(i) != reflect.TypeOf(test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", reflect.TypeOf(i), reflect.TypeOf(test.expected))
			}
		})
	}
}

func TestConvertFlagParamsJsonStringFromFileOrStdin(t *testing.T) {
	tests := map[string]struct {
		param    string
		expected string
	}{
		"if defined input with valid data": {
			param:    "testdata/input_sample.txt",
			expected: "{\"a\": \"1\", \"b\": \"2\", \"c\": \"3\"}",
		},
		"if defined invalid input": {
			param:    "testdata/input_sample_whatever.txt",
			expected: "cannot open",
		},
		"if defined input as STDIN": {
			param:    "-",
			expected: "{\"a\": \"1\", \"b\": \"2\", \"c\": \"3\"}",
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.param == "-" {
				content := []byte("{\"a\": \"1\", \"b\": \"2\", \"c\": \"3\"}")
				tmpfile, err := ioutil.TempFile("", "example")
				if err != nil {
					t.Errorf("Unexpected error: %v\n", err)
				}
				defer os.Remove(tmpfile.Name()) // clean up

				if _, err := tmpfile.Write(content); err != nil {
					t.Errorf("Unexpected error: %v\n", err)
				}
				if _, err := tmpfile.Seek(0, 0); err != nil {
					t.Errorf("Unexpected error: %v\n", err)
				}
				oldStdin := os.Stdin
				defer func() { os.Stdin = oldStdin }() // Restore original Stdin
				os.Stdin = tmpfile
			}

			s, err := ConvertFlagParamsJsonStringFromFileOrStdin(test.param)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && reflect.TypeOf(s) != reflect.TypeOf(test.expected) {
				t.Errorf("Unexpected response: %v # %v. Expected: %v\n", s, reflect.TypeOf(s), reflect.TypeOf(test.expected))
			}
		})
	}
}
