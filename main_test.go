package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestUpdateLocalE2E(t *testing.T) {
	const happyTestFilePath = "testing/testhcl.tf"

	// create test file
	file, err := os.CreateTemp("", "testhcl.tf")
	assert.NoError(t, err, "Error creating test file")

	_, err = file.WriteString(`locals {
		# pin the target versions of the code
		other_code_version = "3.3.3.3"
		code_version       = "1.1.1.1"
	  }

	  output "test_version_string" {
		value = var.other_code_version
	  }

	  output "test_version_number" {
		value = var.code_version
	  }
`)
	assert.NoError(t, err, "Error writing to test file")

	file.Close()

	// set environment variables
	os.Setenv("INPUT_FILEPATH", file.Name())
	os.Setenv("INPUT_VARNAME", "code_version")
	os.Setenv("INPUT_VALUE", "v2.55.4")

	// run e2e
	main()

	// load and check modified file matches expected file
	fileData, err := os.ReadFile(file.Name())
	assert.NoError(t, err, "Error reading test file")

	happyTestFile, err := os.ReadFile(happyTestFilePath)
	assert.NoError(t, err, "Error reading updated test file")
	assert.Equal(t, string(happyTestFile), string(fileData), "Test file and expected file don't match")
}
