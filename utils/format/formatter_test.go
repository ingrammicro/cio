// Copyright (c) 2017-2021 Ingram Micro Inc.

package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultFormatter(t *testing.T) {

	assert := assert.New(t)

	formatter = nil
	f := GetFormatter()
	assert.NotNil(f, "Formatter shouldn't be nil")
}
