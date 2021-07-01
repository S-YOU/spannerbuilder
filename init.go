package spannerbuilder

import (
	"os"
	"strings"
)

var (
	debug bool
)

func init() {
	debug = makeBool(os.Getenv("DB_DEBUG"))
}

func makeBool(s string) bool {
	return strings.ToLower(s) == "true" || s == "1"
}
