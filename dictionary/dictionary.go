package dictionary

const (
	errNotFound = DictionaryErr("could not find the word you were looking for")
	errKeyExists = DictionaryErr("key already exists")
	errWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
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

func (d Dictionary) Update(key, val string) error {
	_, err := d.Search(key)

	switch err {
	case errNotFound:
		return errWordDoesNotExist 
	case nil:
		d[key] = val
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
