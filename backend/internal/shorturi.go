package internal

import "strings"

const validChars = "bcdfghmnprstvz23456789"

func ShortURIfromID(id uint) string {
	uri := ""
	for id > 0 {
		uri = string(validChars[id%uint(len(validChars))]) + uri
		id = id / uint(len(validChars))
	}
	return uri
}

func IDFromShortURI(uri string) uint {
	id := uint(0)
	for _, c := range uri {
		if !strings.Contains(validChars, string(c)) {
			return 0
		}
		id = id*uint(len(validChars)) + uint(strings.Index(validChars, string(c)))
	}
	return id
}
