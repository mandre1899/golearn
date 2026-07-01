package dictionary

type Dictionary map[string]string

func Search(m Dictionary, toSearch string) string {
	return m[toSearch]
}
