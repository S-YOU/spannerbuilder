package spannerbuilder

type Builder struct {
	table  string
	sel    string
	group  string
	sample string
	index  string
	cols   []string
	keys   []string
	wheres []string
	orders []string
	joins  []string
	having []string
	limit  int
	args   map[string]interface{}
}

func NewSpannerBuilder(table string, cols, keys []string) *Builder {
	return &Builder{
		table: table,
		cols:  cols,
		keys:  keys,
		args:  make(map[string]interface{}),
	}
}
