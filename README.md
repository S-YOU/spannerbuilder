# spannerbuilder
spanner sql builder for select statements

### supported methods
- Select
    - `.Select("field_name")`
    - `.Select("DISTINCT field_name", "field_name")`
    - `.Select("COUNT(*)", "field_name")`
- Join
    - `.Join("table_name USING(field_name)")`
    - `.Join("table_name USING(field_name)", "LEFT OUTER")`
    - `.Join("table_name USING(field_name)", "RIGHT", "OUTER")`
    - `.Join("table_name USING(field_name)").Join("table_name USING(field_name)")`
- OrderBy
- Where
    - `.Where("field_name <op> ?", varName)`
    - `.Where("field_name <op> @field_name", Params{"field_name": verName})`
- Limit
