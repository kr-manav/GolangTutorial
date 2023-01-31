# GolangTutorial

### Function Definition
#### Syntax : func [name](var_name datatype) return_datatype {}

```go
func add(x int, y int) int {
	return x + y
}
```

#### For multiple returns, Syntax: func [name](var_name datatype) (return_datatype1, return_datatype2) {}

```go
func swap(x, y string) (string, string) {
	return y, x
}
```
### Variables Definition
#### Syntax : 1) var i int = 0;var j int = 0, 2) i, j:=0, 0 3) var i, j int = 0

```go
var i int = 0
var j int = 0
var a, b int = 0
c,d := 0, 0
```
