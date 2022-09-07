package maps

type Dictionary map[string]string

func (d Dictionary) Search(term string) string {
	if meaning, ok := d[term]; ok {
		return meaning
	}
	return ""
}
