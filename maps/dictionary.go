package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(term string) (string, error) {
	if definition, ok := d[term]; ok {
		return definition, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
