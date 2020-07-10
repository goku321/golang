package main

import "strings"

func joinString(args ...string) string {
	return strings.Join(args, "|")
}

func joinNString() string {
	return strings.Join([]string{"a", "b", "c", "de"}, "|")
}
