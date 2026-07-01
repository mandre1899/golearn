package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(toSearch string) (string, error) {
	val, ok := d[toSearch]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return val, nil
}
