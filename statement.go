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

	for _, x := range b.joins {
		s.WriteString(x)
	}

	if len(b.wheres) > 0 {
		s.WriteString(" WHERE ")
		s.WriteString(strings.Join(b.wheres, " AND "))
	}

	if len(b.orders) > 0 {
		s.WriteString(" ORDER BY ")
		s.WriteString(strings.Join(b.orders, ", "))
	}

	if b.limit > 0 {
		s.WriteString(" LIMIT ")
		s.WriteString(strconv.Itoa(b.limit))
	}

	return spanner.Statement{SQL: s.String(), Params: b.args}
}
