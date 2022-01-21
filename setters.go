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

func (b *Builder) Index(index string) *Builder {
	b.index = index
	return b
}

func (b *Builder) Statement(sql string, args ...interface{}) *Builder {
	var target []string
	b.updateArgs(sql, args, &target, true)
	b.sql = target[0]
	return b
}

func (b *Builder) Select(s string, cols ...string) *Builder {
	// check for backward compatibility
	if len(cols) == 0 {
		b.sel = s
		if strings.IndexByte(s, ',') < 0 {
			b.cols = []string{s} // simple query with one field selected `.Select("field_name")`
		} else {
			b.cols = []string{} // complex query, no column names will be stored in `cols` variable
		}
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
	b.updateArgs(s, args, &b.wheres, false)
	return b
}

func (b *Builder) GroupBy(s string) *Builder {
	b.group = s
	return b
}

func (b *Builder) Having(s string, args ...interface{}) *Builder {
	b.updateArgs(s, args, &b.having, false)
	return b
}

func (b *Builder) TableSample(s string) *Builder {
	b.sample = s
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

func (b *Builder) Offset(i int) *Builder {
	b.offset = i
	return b
}

func (b *Builder) updateArgs(s string, args []interface{}, target *[]string, inline bool) {
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
		si := strconv.Itoa(i + xargs)
		k := "_arg" + si
		s = strings.Replace(s, "?", "@"+k, 1)
		if inline && strings.Contains(s, "{"+si+"}") {
			s = strings.Replace(s, "{"+si+"}", fmt.Sprint(args[i]), -1)
		}
		b.args[k] = args[i]
	}
	*target = append(*target, s)
}
