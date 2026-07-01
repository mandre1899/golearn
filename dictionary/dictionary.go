package dictionary

import (
	"errors"
)

var errNotFound = errors.New("could not find the word you were looking for")
var errKeyExists = errors.New("key already exists")

type Dictionary map[string]string

func (d Dictionary) Search(toSearch string) (string, error) {
	val, ok := d[toSearch]
	if !ok {
		return "", errNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, val string) error {
	_, ok := d[key]
	if ok {
		return errKeyExists
	}
	d[key] = val
	return nil
}
