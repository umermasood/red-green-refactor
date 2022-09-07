package maps

func Search(dictionary map[string]string, term string) string {
	if meaning, ok := dictionary[term]; ok {
		return meaning
	}
	return ""
}
