package spannerbuilder

type Builder struct {
	sql    string
	table  string
	sel    string
	group  string
	sample string
	index  string
	cols   []string
	froms  []string
	keys   []string
	wheres []string
	orders []string
	joins  []string
	having []string
	unions []string
	limit  int
	offset int
	uLim   int
	uOfs   int
	uOdrs  []string
	args   map[string]interface{}
}

func NewSpannerBuilder(table string, cols, keys []string, args ...map[string]interface{}) *Builder {
	bArgs := make(map[string]interface{})
	for _, arg := range args {
		for k, v := range arg {
			bArgs[k] = v
		}
	}
	return &Builder{
		table: table,
		cols:  cols,
		keys:  keys,
		args:  bArgs,
	}
}
