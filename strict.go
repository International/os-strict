package strict

import (
	"errors"
	"fmt"
	"os"
)

// MustGetenv returns the value of variable if it's not blank, or panics
func MustGetenv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("variable %s is not set", name))
	}
	return value
}

// Getenv returns the value of variable if it's not blank, or a blank string and an error
func Getenv(name string) (string, error) {
	value := os.Getenv(name)
	if value == "" {
		return "", errors.New(fmt.Sprintf("variable %s is not set", name))
	}
	return value, nil
}

// GetenvWithDefault returns the value of variable if it's not blank, or defaultValue
func GetenvWithDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}
