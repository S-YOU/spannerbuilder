package spannerbuilder

import (
	"fmt"
	"strconv"
	"strings"
)

func (b *Builder) From(s string, args ...interface{}) *Builder {
	if s != "" {
		b.updateArgs(s, args, &b.froms, defaultWhiteList)
	}
	return b
}

func (b *Builder) Index(index string) *Builder {
	b.index = index
	return b
}

func (b *Builder) Statement(sql string, args ...interface{}) *Builder {
	b.sql = b.updateArgs(sql, args, nil, nil)
	return b
}

func (b *Builder) Select(s string, args ...interface{}) *Builder {
	b.updateArgs(s, args, &b.sels, defaultWhiteList)
	return b
}

func (b *Builder) SetColumns(cols ...string) *Builder {
	b.cols = cols
	return b
}

func (b *Builder) Join(s string, args ...interface{}) *Builder {
	var join string
	if len(args) == 0 {
		join = fmt.Sprintf(" INNER JOIN %s", s)
	} else if joinType, ok := args[0].(string); ok && validJoins[joinType] {
		join = fmt.Sprintf(" %s JOIN %s", joinType, s)
		args = args[1:]
	} else {
		join = fmt.Sprintf(" INNER JOIN %s", s)
	}
	b.updateArgs(join, args, &b.joins, defaultWhiteList)
	return b
}

func (b *Builder) Where(s string, args ...interface{}) *Builder {
	if s != "" {
		b.updateArgs(s, args, &b.wheres, defaultWhiteList)
	}
	return b
}

func (b *Builder) GroupBy(s string) *Builder {
	b.group = s
	return b
}

func (b *Builder) Having(s string, args ...interface{}) *Builder {
	if s != "" {
		b.updateArgs(s, args, &b.having, defaultWhiteList)
	}
	return b
}

func (b *Builder) TableSample(s string) *Builder {
	b.sample = s
	return b
}

func (b *Builder) OrderBy(s string, args ...interface{}) *Builder {
	if s != "" {
		if len(b.unions) > 0 {
			b.updateArgs(s, args, &b.uOdrs, defaultWhiteList)
		} else {
			b.updateArgs(s, args, &b.orders, defaultWhiteList)
		}
	}
	return b
}

func (b *Builder) Limit(i int) *Builder {
	if len(b.unions) > 0 {
		b.uLim = i
	} else {
		b.limit = i
	}
	return b
}

func (b *Builder) Offset(i int) *Builder {
	if len(b.unions) > 0 {
		b.uOfs = i
	} else {
		b.offset = i
	}
	return b
}

func (b *Builder) Union(sel Selector, unionTypes ...string) *Builder {
	stmt := sel.GetSelectStatement()
	for k, v := range stmt.Params {
		if _, ok := b.args[k]; ok {
			newKey := "_arg" + strconv.Itoa(len(b.args))
			stmt.SQL = strings.Replace(stmt.SQL, k, newKey, -1)
			b.args[newKey] = v
		} else {
			b.args[k] = v
		}
	}
	unionType := "ALL"
	if len(unionTypes) > 0 {
		unionType = unionTypes[0]
	}
	b.unions = append(b.unions, fmt.Sprintf("UNION %s\n(%s)", unionType, stmt.SQL))
	return b
}
