package main

import "strings"

func joinString(args ...string) string {
	return strings.Join(args, "|")
}
