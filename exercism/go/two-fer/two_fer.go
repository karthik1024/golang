package twofer

// ShareWith returns a string based on if a name is given  or not.
func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	} else {
		return "One for you, one for me."
	}
}
