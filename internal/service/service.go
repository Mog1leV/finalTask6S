package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorseOrNot(s string) bool {
	for _, r := range s {
		if !(r == '-' || r == '.' || r == ' ') {
			return false
		}
	}
	return true
}

func ConverterTo(s string) string {
	if isMorseOrNot(s) {
		s = morse.ToText(s)
	} else {
		s = morse.ToMorse(s)
	}
	return s
}
