package spannerbuilder

import (
	"fmt"
	"strconv"
	"strings"
)

func (b *Builder) From(table string) *Builder {
	b.table = table
	return b
}

func (b *Builder) Select(s string, cols ...string) *Builder {
	if len(cols) == 0 {
		b.cols = strings.Split(s, ",")
	} else {
		b.sel = s
		b.cols = cols
	}
	return b
}

func (b *Builder) Join(s string, joinType ...string) *Builder {
	if len(joinType) == 0 {
		b.joins = append(b.joins, fmt.Sprintf(" JOIN %s", s))
	} else {
		b.joins = append(b.joins, fmt.Sprintf(" %s JOIN %s", strings.Join(joinType, " "), s))
	}
	return b
}

func (b *Builder) Where(s string, args ...interface{}) *Builder {
	b.updateArgs(s, args, &b.wheres)
	return b
}

func (b *Builder) GroupBy(s string) *Builder {
	b.group = s
	return b
}

func (b *Builder) Having(s string, args ...interface{}) *Builder {
	b.updateArgs(s, args, &b.having)
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

func (b *Builder) updateArgs(s string, args []interface{}, target *[]string) {
	if len(args) == 1 {
		if m, ok := args[0].(map[string]interface{}); ok {
			for k, v := range m {
				b.args[k] = v
			}
			*target = append(*target, s)
			return
		}
	}
	xargs := len(b.args)
	for i := 0; i < len(args); i++ {
		k := "_arg" + strconv.Itoa(i+xargs)
		s = strings.Replace(s, "?", "@"+k, 1)
		b.args[k] = args[i]
	}
	*target = append(*target, s)
}
