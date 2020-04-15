package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"os"
	"reflect"
)

// FlagConvertParams converts cli parameters in API callable params
func FlagConvertParams(c *cli.Context) *map[string]interface{} {
	v := make(map[string]interface{})
	for _, flag := range c.FlagNames() {
		if c.IsSet(flag) {
			v[flag] = c.String(flag)
		}
	}
	return &v
}

// FlagConvertParamsJSON converts cli parameters in API callable params, and encodes JSON parameters
func FlagConvertParamsJSON(c *cli.Context, jsonFlags []string) (*map[string]interface{}, error) {
	v := make(map[string]interface{})
	for _, flag := range c.FlagNames() {
		if c.IsSet(flag) {

			// check if field is json
			isJSON := false
			if jsonFlags != nil {
				for _, js := range jsonFlags {
					if js == flag {
						isJSON = true
						break
					}
				}
			}

			if isJSON {
				// parse json before assigning to map
				var p interface{}
				err := json.Unmarshal([]byte(c.String(flag)), &p)
				if err != nil {
					return nil, fmt.Errorf("flag %s isn't a valid JSON. %s", flag, err)
				}
				v[flag] = p
			} else {
				v[flag] = c.String(flag)
			}
		}
	}
	return &v, nil
}

// ItemConvertParams converts API items into map of interface
func ItemConvertParams(item interface{}) (*map[string]interface{}, error) {

	it := reflect.ValueOf(item)
	nf := it.NumField()
	v := make(map[string]interface{})

	for i := 0; i < nf; i++ {
		v[it.Type().Field(i).Name] = fmt.Sprintf("%s", it.Field(i).Interface())
		// if value, ok :=  it.Field(i).Interface().(string); ok {
		// 	v[it.Type().Field(i).Name] = value
		// } else {
		// 	return nil, fmt.Errorf("Interface couldn't be converted to map of strings. Field: %s", it.Type().Field(i).Name)
		// }
	}
	return &v, nil
}

// ConvertFlagParamsJsonFromFileOrStdin returns the json representation of parameters taken from the input file or STDIN
func ConvertFlagParamsJsonFromFileOrStdin(c *cli.Context, dataIn string) (map[string]interface{}, error) {
	var reader io.Reader
	var content map[string]interface{}
	if dataIn == "-" {
		reader = os.Stdin
		dataIn = "STDIN"
	} else {
		jsonFile, err := os.Open(dataIn)
		if err != nil {
			return nil, fmt.Errorf("cannot open %s to read JSON params: %v", dataIn, err)
		}
		defer jsonFile.Close()
		reader = jsonFile
	}
	if err := json.NewDecoder(reader).Decode(&content); err != nil {
		return nil, fmt.Errorf("cannot read JSON params from %s: %v", dataIn, err)
	}
	return content, nil
}

// ConvertFlagParamsJsonStringFromFileOrStdin returns the json string representation of parameters taken from the input file or STDIN
func ConvertFlagParamsJsonStringFromFileOrStdin(c *cli.Context, dataIn string) (string, error) {
	var reader io.Reader

	if dataIn == "-" {
		reader = os.Stdin
		dataIn = "STDIN"
	} else {
		jsonFile, err := os.Open(dataIn)
		if err != nil {
			return "", fmt.Errorf("cannot open %s to read params: %v", dataIn, err)
		}
		defer jsonFile.Close()
		reader = jsonFile
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return "", fmt.Errorf("cannot read %s: %v", dataIn, err)
	}
	return buf.String(), nil
}
