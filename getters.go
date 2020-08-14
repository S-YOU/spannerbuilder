package spannerbuilder

import (
	"strings"
)

func (b *Builder) Columns() []string {
	cols := make([]string, 0, len(b.cols))
	for _, x := range b.cols {
		if i := strings.Index(x, " AS "); i != -1 {
			x = x[i+4:]
		}
		x = strings.TrimSpace(x)
		cols = append(cols, x)
	}
	return cols
}
