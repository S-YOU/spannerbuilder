package spannerbuilder

import (
	"fmt"
	"strings"
)

func (b *Builder) From(s string, args ...interface{}) *Builder {
	if s != "" {
		b.updateArgs(s, args, &b.froms, nil)
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

func (b *Builder) Select(s string, cols ...string) *Builder {
	// check for backward compatibility
	if len(cols) == 0 {
		b.sel = s
		if strings.IndexByte(s, ',') < 0 && strings.IndexByte(s, '.') < 0 {
			b.cols = []string{s} // simple query with one field selected `.Select("field_name")`
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

func (b *Builder) Union(sel Selector, unionType ...string) *Builder {
	sql := sel.GetSelectStatement().SQL
	if len(unionType) == 0 {
		b.unions = append(b.unions, fmt.Sprintf("UNION ALL\n(%s)", sql))
	} else {
		b.unions = append(b.unions, fmt.Sprintf("UNION %s\n(%s)", unionType[0], sql))
	}
	return b
}
