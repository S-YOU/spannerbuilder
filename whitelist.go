package spannerbuilder

var (
	defaultWhiteList = map[string]bool{
		"=":    true,
		"<":    true,
		">":    true,
		"<=":   true,
		">=":   true,
		"!=":   true,
		"<>":   true,
		"IN":   true,
		"LIKE": true,
		"ASC":  true,
		"DESC": true,
		"":     true,
	}
)
