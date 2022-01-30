# spannerbuilder
spanner sql builder for select statements

### supported methods
- Select
    - Syntax: `.Select(statement, col1, col2, ...)`
    - `.Select("field_name")` // special case, same as `.Select("field_name", "field_name")`
    - `.Select("DISTINCT field_name", "field_name")`
    - `.Select("COUNT(*)", "field_name")`
    - `.Select("field_name1, field_name2", "field_name1", "field_name2")`
    - `.Select("field_name1").Select("field_name2")` // only `field_name2` will be selected
- Join
    - `.Join("table_name USING(field_name)")`
    - `.Join("table_name USING(field_name)", "LEFT")`
    - `.Join("table_name USING(field_name)", "RIGHT OUTER")`
    - `.Join("table_name USING(field_name)").Join("table_name USING(field_name)")`
    - `.Join("table_name1@{FORCE_INDEX=index_name} ON table_name1.field_name = table_name2.field_name")`
- Index
    - `.Index("index_name")`
- OrderBy
    - `.OrderBy("field_name")`
    - `.OrderBy("field_name DESC")`
- Where
    - `.Where("field_name <op> ?", varName)`
    - `.Where("field_name <op> @field_name", M{"field_name": varName})`
- GroupBy
    - `.Select("Sum(field_name1), field_name2", "field_name1", "field_name2").
            GroupBy("group_by_expression")`
- Having (with GroupBy, Select)
    - `.Select("Sum(field_name1), field_name2", "field_name1", "field_name2").
            GroupBy("group_by_expression").Having("SUM(field_name1) > 0")`
    - `.Select("Sum(field_name1), field_name2", "field_name1", "field_name2").
            GroupBy("group_by_expression").Having("SUM(field_name1) > ?", 1)`
    - `.Select("Sum(field_name1), field_name2", "field_name1", "field_name2").
            GroupBy("group_by_expression").Having("SUM(field_name1) > @param0", M{"param0": 1})`
    - Having (with GroupBy, Select, Where)
        - `.Select("Sum(field_name1), field_name2", "field_name1", "field_name2").
                Where("field_name2 == ?", "something").
                GroupBy("group_by_expression").Having("SUM(field_name1) > @param0", M{"param0": 1})`
- Limit
    - `.Limit(20)`
- Offset (with Limit)
    - `.Limit(20).Offset(10)`
- From
    - `.From("table_name")`
    - `.From("table_name@{FORCE_INDEX=index_name}")`
    - `.From("UNNEST(GENERATE_ARRAY(0, @fieldName)) AS field_name", M{"fieldName": varName})`
- TableSample
    - `.TableSample("RESERVOIR (100 ROWS)")`,
    - `.TableSample("BERNOULLI (0.1 PERCENT)")`
- Statement
    - `.Statement("SELECT * FROM users WHERE user_id = ?", "test")`
    - `.Statement("SELECT * FROM users WHERE user_id = '{0}'", "test")`
    - `.Statement("SELECT * FROM users WHERE user_id IN ({0}, {1})", 1, 2)`
    - `.Statement("SELECT * FROM users WHERE user_id = @userId", M{"userId": "test"})`
    - `.Statement("SELECT * FROM users WHERE user_id = '{userId}'", M{"userId": "test"})`
- SubQuery
  - `XX(M{"subQuery": YY()}).From("({subQuery})")`
- Union
    - `XX.From("users").Union(YY.From("admins"))`

### SQL Logging
- set `DB_DEBUG=1` to output SQL Log
