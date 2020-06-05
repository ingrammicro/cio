// Copyright (c) 2017-2022 Ingram Micro Inc.

package utils

import (
	"context"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// Untar decompresses the source file to target file
func Untar(ctx context.Context, source, target string) error {
	if err := os.MkdirAll(target, 0600); err != nil {
		return err
	}

	tarExecutable := "tar"
	if runtime.GOOS == "windows" {
		tarExecutable = "C:\\opscode\\chef\\bin\\tar.exe"
	}
	cmd := exec.CommandContext(ctx, tarExecutable, "-xzf", source, "-C", target)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// FileExists checks file existence
func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// RandomString generates a random string from lowercase letters and numbers
func RandomString(strlen int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

// RemoveDuplicates returns the slice removing duplicates if exist
func RemoveDuplicates(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	var result []string
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

// Contains evaluates whether s contains x.
func Contains(s []string, x string) bool {
	for _, n := range s {
		if x == n {
			return true
		}
	}
	return false
}
