package dictionary

import (
	"errors"
)

var errNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(toSearch string) (string, error) {
	val, ok := d[toSearch]
	if !ok {
		return "", errNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, val string) {
	d[key] = val
}
