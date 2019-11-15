package main

import "errors"

type StringSlice []string

// TODO : Move util methods into its own library

func (s StringSlice) EntryExists(entry string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == entry {
			return true
		}
	}
	return false
}

// NewError creates and returns a new error and returns it
func NewError(errStr string) error {
	return errors.New(errStr)
}
