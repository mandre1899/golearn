package dictionary

type Dictionary map[string]string

func (d Dictionary) Search(toSearch string) string {
	return d[toSearch]
}
