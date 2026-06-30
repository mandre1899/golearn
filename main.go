package main

import (
	"fmt"
	"strings"
)

const (
	spanish = "spanish"
	french = "french"
	englishHelloPrefix = "Hello,"
	spanishHelloPrefix = "Hola,"
	frenchHelloPrefix = "Bonjour,"
)

func Hello(name, lang string) string {
	prefix := greetingPrefix(lang)
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s %s", prefix, name)
}

func greetingPrefix(lang string) (prefix string) {
	switch strings.ToLower(lang) {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
