package main

import (
	"errors"
	"strconv"
	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) string
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

type stringservice struct{}

func (stringservice) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringservice) Count(s string) string {
	return strconv.Itoa(len(s))
}
