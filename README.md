# spannerbuilder
spanner sql builder for select statements

### supported methods
- Select
- Join (just JOIN)
- OrderBy
- Where
    - `.Where("field_name <op> ?", varName)`
    - `.Where("field_name <op> @field_name", Params{"field_name": verName})`
- Limit
