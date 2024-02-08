package pygo

import "strings"

type Str string

// Converts the first character to upper case
func (s Str) Capitalize() Str {
	var ns Str
	if len(s) > 1 {
		ns = Str(strings.ToUpper(string(s[0]))) + s[1:len(s)-1]
	} else if len(s) == 1 {
		ns = Str(strings.ToUpper(string(s[0])))
	} else {
		ns = ""
	}
	return Str(ns)
}

// Converts string into lower case
func (s Str) Casefold() Str {
	return Str(strings.ToLower(string(s)))
}

// Returns a centered string
func (s Str) Center(space int) Str {
	if len(s) >= space {
		return s
	}
	c := space - len(s)
	if c%2 == 0 {
		return Str(strings.Repeat(" ", c/2)) + s + Str(strings.Repeat(" ", c/2))
	}
	c -= 1
	return Str(strings.Repeat(" ", c/2+1)) + s + Str(strings.Repeat(" ", c/2))
}

// Returns the number of times a specified value occurs in a string
func (s Str) Count(value string, i ...int) int {
	var start, end int
	start = 0
	end = len(s)
	if len(i) == 1 {
		start = i[0]
	} else if len(i) > 1 {
		start = i[0]
		end = i[1]
	}
	t := string(s[start:end])
	return strings.Count(t, value)
}
