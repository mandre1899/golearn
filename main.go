package main

import (
	"fmt"
	"strings"
)


const spanish = "spanish"
const french = "french"
const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"

func Hello(name, lang string) string {
	var prefix string
	switch strings.ToLower(lang) {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s %s", prefix, name)
}
