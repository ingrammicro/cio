// Copyright (c) 2017-2021 Ingram Micro Inc.

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

// FlagConvertParamsJSON converts cli parameters in API callable params, and encodes JSON parameters
func FlagConvertParamsJSON(flagName string) (*map[string]interface{}, error) {
	v := make(map[string]interface{})
	if viper.IsSet(flagName) {
		var p interface{}
		err := json.Unmarshal([]byte(viper.GetString(flagName)), &p)
		if err != nil {
			return nil, fmt.Errorf("flag %s isn't a valid JSON. %s", flagName, err)
		}
		v[flagName] = p
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
		// 	return nil, fmt.Errorf("Interface couldn't be converted to map of strings. Field: %s",
		// it.Type().Field(i).Name)
		// }
	}
	return &v, nil
}

// ConvertFlagParamsJsonFromFileOrStdin returns the json representation of parameters taken from the input file or STDIN
func ConvertFlagParamsJsonFromFileOrStdin(dataIn string) (map[string]interface{}, error) {
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

// ConvertFlagParamsJsonStringFromFileOrStdin returns the json string representation of parameters taken from the input
// file or STDIN
func ConvertFlagParamsJsonStringFromFileOrStdin(dataIn string) (string, error) {
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
