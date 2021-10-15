## Summary

- Variable declaration
  - var foo int
  - var foo int = 42
  - foo := 42

- Cant redeclare variables, but can shadow then
- All variables must be used
- Visibility
  - lower case first letter for package scope
  - upper case first letter to export
  - no private scope
- Naming conventions
  - Pascal or camelCase
  - As short as resonable
    - longer names for longer lives
- Type conversions
  - destinationType (variable)
  - use strconv package for strings
