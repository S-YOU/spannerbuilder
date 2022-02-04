package spannerbuilder

import (
	"regexp"
	"strings"
)

var (
	selectExprAsteriskRegexp = regexp.MustCompile(`(?i)COUNT\((?:\w+\.)?\*\)|(?:(\w+)\.)?\*`)
	fromExprAliasRegexp      = regexp.MustCompile(`(?i)(\w+)\b\s*(?:{[^}]+})?(?:\s*AS)?\s*(\w+)`)
	joinExprAliasRegexp      = regexp.MustCompile(`(?i)JOIN\s*(\w+)\b(?:\s*AS)?\s*(\w+)`)
)

func (b *Builder) buildFromFieldMap(fs string) {
	f := fromExprAliasRegexp.FindAllStringSubmatchIndex(fs, -1)
	if len(f) > 0 {
		for _, x := range f {
			if strings.ToUpper(fs[x[2]:x[3]]) == "AS" {
				continue
			}
			table, alias := fs[x[2]:x[3]], ""
			if x[4] >= 0 {
				alias = fs[x[4]:x[5]]
			}
			if fields, ok := b.tblMap[table]; ok {
				b.fldMap[table] = fields
				if alias != "" {
					b.fldMap[alias] = fields
				}
			}
		}
	}
}

func (b *Builder) buildJoinFieldMap(fs string) {
	f := joinExprAliasRegexp.FindAllStringSubmatchIndex(fs, -1)
	if len(f) > 0 {
		for _, x := range f {
			upr := strings.ToUpper(fs[x[4]:x[5]])
			if upr == "ON" || upr == "USING" {
				continue
			}
			table, alias := fs[x[2]:x[3]], ""
			if x[4] >= 0 {
				alias = fs[x[4]:x[5]]
			}
			if fields, ok := b.tblMap[table]; ok {
				b.fldMap[table] = fields
				if alias != "" {
					b.fldMap[alias] = fields
				}
			}
		}
	}
}

func (b *Builder) expandAsterisk(s *string) {
	if s != nil && b.tblMap != nil {
		m := selectExprAsteriskRegexp.FindAllStringSubmatchIndex(*s, -1)
		if len(m) > 0 {
			sb := strings.Builder{}
			lastIndex := 0
			for _, x := range m {
				sb.WriteString((*s)[lastIndex:x[0]])
				matchStr := (*s)[x[0]:x[1]]
				if strings.HasPrefix(matchStr, "COUNT(") || strings.HasPrefix(matchStr, "STRUCT") {
					sb.WriteString(matchStr)
					lastIndex = x[1]
					continue
				}
				table := b.table
				if x[2] >= 0 {
					table = (*s)[x[2]:x[3]]
				}
				if fields, ok := b.fldMap[table]; ok {
					for i, x := range fields {
						if i > 0 {
							sb.WriteByte(',')
						}
						sb.WriteString(kwQuoted(table))
						sb.WriteString(".")
						sb.WriteString(kwQuoted(x))
					}
				} else {
					// fallback for unknown tables
					sb.WriteString(matchStr)
				}
				lastIndex = x[1]
			}

			sb.WriteString((*s)[lastIndex:])
			*s = sb.String()
		}
	}
}

func (b *Builder) SetTableFieldMap(m map[string][]string) *Builder {
	b.tblMap = m
	for k, v := range m {
		b.fldMap[k] = v
	}
	return b
}
