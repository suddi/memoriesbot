package config

import (
	"os"
	"strconv"
	"testing"
)

func TestGetEnv(t *testing.T) {
	errorMessage := "CASE 1: Should be able to retrieve set environment variable"
	key := "TEST_VARIABLE"
	value := "TEST_VALUE"
	defaultValue := "DEFAULT_TEST_VALUE"

	os.Setenv(key, value)

	result := getEnv(key, defaultValue)

	if result != value {
		t.Error(errorMessage)
	}

	errorMessage = "CASE 2: Should be able to default environment variable value"

	os.Unsetenv(key)

	result = getEnv(key, defaultValue)

	if result != defaultValue {
		t.Error(errorMessage)
	}
}

func TestGetEnvAsInt(t *testing.T) {
	errorMessage := "CASE 1: Should be able to retrieve set environment variable"
	key := "TEST_VARIABLE"
	value := "42"
	defaultValue := 0

	os.Setenv(key, value)

	result := getEnvAsInt(key, defaultValue)

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		t.Error(errorMessage)
	}

	if result != valueInt {
		t.Error(errorMessage)
	}

	errorMessage = "CASE 2: Should be able to default environment variable value"

	os.Unsetenv(key)

	result = getEnvAsInt(key, defaultValue)

	if result != defaultValue {
		t.Error(errorMessage)
	}
}

func TestGetEnvAsBool(t *testing.T) {
	errorMessage := "CASE 1: Should be able to retrieve set environment variable"
	key := "TEST_VARIABLE"
	value := "true"
	defaultValue := false

	os.Setenv(key, value)

	result := getEnvAsBool(key, defaultValue)

	valueBool, err := strconv.ParseBool(value)
	if err != nil {
		t.Error(errorMessage)
	}

	if result != valueBool {
		t.Error(errorMessage)
	}

	errorMessage = "CASE 2: Should be able to default environment variable value"

	os.Unsetenv(key)

	result = getEnvAsBool(key, defaultValue)

	if result != defaultValue {
		t.Error(errorMessage)
	}
}
