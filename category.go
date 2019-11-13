package main

import (
	"errors"
	"fmt"
)

type Category struct {
	Values StringSlice
}

// Get get the values of a category enum type
func (c *Category) Get() (StringSlice, error) {
	if len(c.Values) < 1 {
		return nil, errors.New(fmt.Sprintf("category value is a required parameter"))
	}
	e := Enum{
		name: categoryEnumTypeName,
	}
	err := e.Get()
	return e.values, err
}

// Exists checks if a value already exists in the category enum
func (c *Category) Exists() (bool, error) {
	if c.Values[0] == "" {
		return false, errors.New(fmt.Sprintf("category value is a required parameter"))
	}
	e := Enum{
		name: categoryEnumTypeName,
	}
	err := e.Get()
	if err != nil {
		return false, err
	}
	fmt.Println(e.values)
	if e.values.EntryExists(c.Values[0]) {
		return true, errors.New(fmt.Sprintf("category %s already exists", c.Values[0]))
	} else {
		return false, nil
	}
}

// Add adds a new category to the category enum
func (c *Category) Add() error {
	if len(c.Values) < 1 {
		return errors.New(fmt.Sprintf("category value is a required parameter"))
	}
	e := Enum{
		name:   categoryEnumTypeName,
		values: c.Values,
	}
	err := e.Update()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Category) ValidateCategoryValues() error {
	if len(c.Values) < 1 {
		return errors.New(fmt.Sprintf("atlease one category value is required"))
	}
	return nil
}
