package spannerbuilder

import (
	"strconv"
	"strings"
)

func (b *Builder) Select(s string) *Builder {
	b.cols = strings.Split(s, ",")
	return b
}

func (b *Builder) Join(s string) *Builder {
	b.joins = append(b.joins, s)
	return b
}

func (b *Builder) Where(s string, args ...interface{}) *Builder {
	if len(args) == 1 {
		if v, ok := args[0].(map[string]interface{}); ok {
			b.args = v
			b.wheres = append(b.wheres, s)
			return b
		}
	}
	xargs := len(b.args)
	for i := 0; i < len(args); i++ {
		k := "arg" + strconv.Itoa(i+xargs)
		s = strings.Replace(s, "?", "@"+k, 1)
		b.args[k] = args[i]
	}
	b.wheres = append(b.wheres, s)
	return b
}

func (b *Builder) OrderBy(s string) *Builder {
	b.orders = append(b.orders, s)
	return b
}

func (b *Builder) Limit(i int) *Builder {
	b.limit = i
	return b
}
