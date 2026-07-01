package dictionary

const (
	errNotFound = DictionaryErr("could not find the word you were looking for")
	errKeyExists = DictionaryErr("key already exists")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

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
