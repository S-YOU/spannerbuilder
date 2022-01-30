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

	validJoins = map[string]bool{
		"":            true,
		"INNER":       true,
		"OUTER":       true,
		"CROSS":       true,
		"LEFT":        true,
		"RIGHT":       true,
		"FULL":        true,
		"FULL OUTER":  true,
		"LEFT OUTER":  true,
		"RIGHT OUTER": true,
	}
)
