package strict

import (
	"errors"
	"fmt"
	"os"
)

func MustGetenv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("variable %s is not set", name))
	}
	return value
}

func Getenv(name string) (string, error) {
	value := os.Getenv(name)
	if value == "" {
		return "", errors.New(fmt.Sprintf("variable %s is not set", name))
	}
	return value, nil
}
