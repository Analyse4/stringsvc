// StringService provides operations on strings.
package main

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}
