package spannerbuilder

var (
	keywords = map[string]struct{}{
		"groups": {},
	}
)

func kwQuoted(s string) string {
	if _, ok := keywords[s]; ok {
		return "`" + s + "`"
	}
	return s
}
