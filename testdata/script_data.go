// Copyright (c) 2017-2021 Ingram Micro Inc.

package testdata

import "github.com/ingrammicro/cio/api/types"

// GetScriptData loads test data
func GetScriptData() []*types.Script {

	return []*types.Script{
		{
			ID:          "fakeID0",
			Name:        "fakeName0",
			Description: "this is a description for fake Script 0",
			Code: `#!/bin/bash

if [ ! -f /this/file ];
then
  echo "this file must be created" > /this/file
else
  echo "this file exists"
fi`,
			Parameters: []string{},
		},
		{
			ID:          "fakeID1",
			Name:        "fakeName1",
			Description: "this is a description for fake Script 1",
			Code: `#!/bin/bash

echo "received param $PARAM0"
if [ ! -f /this/file ];
then
  echo "this file must be created" > /this/file
else
  echo "this file exists"
fi`,
			Parameters: []string{"PARAM0"},
		},
	}
}
