# GolangTutorial

## Datatypes in Go
```
bool - false (default)

string - "" (default)

int  int8  int16  int32  int64 - 0 (default)
uint uint8 uint16 uint32 uint64 uintptr- 0 (default)

byte // alias for uint8 - 0 (default)

rune // alias for int32 - 0 (default)
     // represents a Unicode code point

float32 float64 - 0.00 (default)

complex64 complex128 - 0 (default)
```


## Function Definition
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


## Variables Definition
#### Syntax : 1) var i int = 0;var j int = 0, 2) i, j:=0, 0 3) var i, j int = 0

```go
var i int = 0
var j int = 0
var a, b int = 0
c,d := 0, 0
```


## Type Conversions

#### Syntax : new_datatype(variable_name)
```go
i:=10
f:=float(i)
fmt.Println(f) // Output: 10.00
u := uint(f)
fmt.Println(u) // Output: 10
```


## Constants
#### Syntax: const variable_name = value

```go
const i=10
fmt.Println(i) // Output:10
```

### Points to remeber about const : 
#### 1) You cannot change the value of constant variable.
#### 2) Constant cannot be defined and initialized using := .


## Loops

#### Loops in go are similar to C only the parenthesis of for loop is not used
#### Syntax: for init;condition;post {}

```go
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

#### Also, the init and post statements are optional

```go
func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```

#### for loop without init and post statements are while loop.

```go
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```


#### To make the for loop infinite, omit the loop condition
```go
func main() {
	sum := 1
	for {
		sum += sum
	}
	fmt.Println(sum)
}
```


## Decision Making (IF-ELSE)

### IF

#### If statements in Go are similar to that in C, only the parenthesis is not used
#### Syntax: if condition {}

```go
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```

#### We can use a short statement with if condition 

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
```

### IF-ELSE

#### If else is similar to that in C without parenthesis. You can use short statements with if-else in go

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
```

### SWITCH

#### Switch is similar to that in C. You can use short statement with switch in C

```go 
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

#### Switch without condition searches for switch true

```go
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

## Defer Statements

#### Defer statements are used to implement the function or functionality that needs to be implemented whether the function runs successfully or not
#### Defer statements are generally used to ensure that the files are closed when their need is over, or to close the channel, or to catch the panics in the program.

```go
func main() {
	defer fmt.Println("world")
	fmt.Println("Hi")
	fmt.Println("hello")
}
```

#### In Go language, multiple defer statements are allowed in the same program and they are executed in LIFO.

```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) 
	}

	fmt.Println("done") // Output: 9 8 7 6 5 4 3 2 1 0 done
}
```


## Pointers

#### Pointers hold the memory address of that variable or value
#### Pointers have 3 value: 1) address 2) value 3) pointer to that variable
#### for eg : v:=10; p:=&v
#### 1) Address = &p 2) Value = *p 3) Pointer to that variable: p

```go
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

## Structs

#### Struct is a collection of fields
#### Syntax: type Struct_Name struct { field1_Name datatype, field2_Name datatype }

```go
type Vertex struct {
	X int
	Y int
}
```

#### You can have the initialize the struct values as syntax: Struct_name{value1, value2} or Struct_name(field1:value1, field2:value2}
#### Value that are not initialized is made default based on their datatype

```go
v1 = Vertex{1, 2}  // has type Vertex
v2 = Vertex{X: 1}  // Y:0 is implicit
v3 = Vertex{}      // X:0 and Y:0
```
#### Accessing the values of struct as Syntax: instance_of_struct.field1

```go
func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) // Output : 4
}
```
 ### OR
 
 ```go
 func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

## Arrays

#### Arrays are group or list of values having same datatypes
#### Syntax: var array_name = [size(optional]datatype{values}

```go
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // Output : Hello World
	fmt.Println(a) // Output: [Hello World]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // Output: [2 3 5 7 11 13]
}
```

### Slices

#### Slices are used to take the part of the array and store it in different variable 
#### SLices have 3 properties: 1) Pointers 2) Length 3) Capacity
#### Syntax: new_variable := old_variable[lower_limit:Upper_limit]
#### Here upper limit is exclusive whereas lowerlimit is inclusive.

```go
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s) // 3 5 7
}
```
#### Slices are based more on capacity. For below example, if you ask for the slice more than the length it will take the value from original array according to his capacity

```go
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // Output: len=6 cap=6 [2 3 5 7 11 13] 

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // Output: len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // Output: len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // Output: len2= cap=4 [5 7]
	
	s = s[1:4]
	printSlice(s) // Output: len=4 cap=3 [5 7 11]
	
	s = s[1:5]
	printSlice(s) // Output: error as it asked for 4 numbers but capacity is 3
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

### Creating a slice with make function
#### make function helps to create the array wthout any error
#### Syntax: make([] datatype)
	
