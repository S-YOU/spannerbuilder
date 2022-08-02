package spannerbuilder

import (
	"fmt"
	"strconv"
	"strings"
)

const maxNested = 5

func (b *Builder) Params(args map[string]interface{}) *Builder {
	for k, v := range args {
		b.args[k] = v
	}
	return b
}

func (b *Builder) updateArgs(s string, args []interface{}, target *[]string, whitelist map[string]bool) string {
	start := 0
	if len(args) == 1 {
		if m, ok := args[0].(map[string]interface{}); ok {
			for k, v := range m {
				b.args[k] = v
			}
			start = 1
		}
	}
	xargs := len(b.args)
	for i := start; i < len(args); i++ {
		ii := strconv.Itoa(i)
		k := "_arg" + strconv.Itoa(i+xargs)
		s = strings.Replace(s, "?", "@"+k, 1)
		b.args[k] = args[i]
		b.replaceArg(&s, "{"+ii+"}", args[i], whitelist)
	}
	found := true
	for i := 0; i < maxNested && found; i++ {
		found = false
		for k, v := range b.args {
			if b.replaceArg(&s, "{"+k+"}", v, whitelist) {
				found = true
			}
		}
	}
	if target != nil && s != "" {
		*target = append(*target, s)
	}
	return s
}

func (b *Builder) replaceArg(s *string, old string, n interface{}, whitelist map[string]bool) bool {
	if strings.Contains(*s, old) {
		if x, ok := n.(Selector); ok {
			stmt := x.GetSelectStatement()
			for k, v := range stmt.Params {
				b.args[k] = v
			}
			*s = strings.Replace(*s, old, "("+stmt.SQL+")", -1)
			return true
		} else {
			repl := fmt.Sprint(n)
			if whitelist == nil || whitelist[repl] {
				*s = strings.Replace(*s, old, fmt.Sprint(n), -1)
				return true
			}
		}
	}
	return false
}
