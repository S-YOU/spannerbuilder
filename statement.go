package spannerbuilder

import (
	"strconv"
	"strings"

	"cloud.google.com/go/spanner"
)

func (b *Builder) GetSelectStatement() spanner.Statement {
	var s strings.Builder

	s.WriteString("SELECT ")
	if b.sel == "" {
		for i, x := range b.cols {
			if i > 0 {
				s.WriteByte(',')
			}
			if len(b.joins) > 0 {
				s.WriteString(kwQuoted(b.table) + "." + kwQuoted(x))
			} else {
				s.WriteString(kwQuoted(x))
			}
		}
	} else {
		s.WriteString(b.sel)
	}

	s.WriteString(" FROM ")
	s.WriteString(kwQuoted(b.table))

	if len(b.index) > 0 {
		s.WriteString("@{FORCE_INDEX=")
		s.WriteString(b.index)
		s.WriteByte('}')
	}

	for _, x := range b.joins {
		s.WriteString(x)
	}

	if len(b.wheres) > 0 {
		s.WriteString(" WHERE ")
		s.WriteString(strings.Join(b.wheres, " AND "))
	}

	if len(b.group) > 0 {
		s.WriteString(" GROUP BY ")
		s.WriteString(b.group)
	}

	if len(b.having) > 0 {
		s.WriteString(" HAVING ")
		s.WriteString(strings.Join(b.having, " AND "))
	}

	if len(b.sample) > 0 {
		s.WriteString(" TABLESAMPLE ")
		s.WriteString(b.sample)
	}

	if len(b.orders) > 0 {
		s.WriteString(" ORDER BY ")
		s.WriteString(strings.Join(b.orders, ", "))
	}

	if b.limit > 0 {
		s.WriteString(" LIMIT ")
		s.WriteString(strconv.Itoa(b.limit))
	}

	if b.offset > 0 {
		s.WriteString(" OFFSET ")
		s.WriteString(strconv.Itoa(b.offset))
	}

	return spanner.Statement{SQL: s.String(), Params: b.args}
}
