# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com
- https://www.linkedin.com/in/tkmagesh

## Schedule
| What? | When? |
| ----- | ----- |
| Commence | 9:00 AM |
| Tea Break | 10:30 AM (15 mins) |
| Lunch Break | 12:30 PM (1 hr) |
| Tea Break | 3:00 PM (15 mins) |
| Wind up   | 5:00 PM |

## Methodology
- No powerpoint
- Code & Discussion
- No dedicated time for Q & A

## Repository
- https://github.com/tkmagesh/cisco-go-nov-2024

## Software Requirements
- Go tools (https://go.dev/dl)
    - Verification
        ```shell
        go version
        ```
- Visual Studio Code (or any other editor)
- Go Extension for VS Code (https://marketplace.visualstudio.com/items?itemName=golang.Go)

## Why Go?
- Simplicity
    - ONLY 25 keywords
    - No access modifiers (public, private, protected etc)
    - No reference types (Everything is a value)
    - No pointer arithmatic support in the language 
    - No classes (only structs)
    - No inheritance (only composition)
    - No exceptions (only errors)
    - No try..catch..finally construct
    - No implicit type conversions
- Managed Concurrency
    - lightweight concurrency model (go routines) - 4KB
    - Built in scheduler
    - Concurrency support is built in the language itself
        - "go" keyword, "chan" data type, channel operator (<-), range construct, select-case constructs
    - Standart Library support
        - "sync" package
        - "sync/atomic" package
- Close to hardware
    - No VM, No JIT
    - Less startup time
    - One build per deployment platform
    - Builtin support for cross compilation

## Standard Library Docs 
- https://pkg.go.dev/std 

## Compilation
```shell
go build <file_name.go>
```

## Compile & Execute
```shell
go run <file_name.go>
```

## Get the list of suppported platforms
```shell
go tool dist list
```

## Get the env variables & values
```shell
go env
```
```shell
go env <var_1> <var_2>
```

## Env variables for cross compilation
- GOOS
- GOARCH

## Cross compilation
```shell
GOOS=windows GOARCH=amd64 go build 01-hello-world.go
```

## Data Types
- string
- bool
- integers
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integers
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating points
    - float32
    - float64
- complex
    - complex64 ( real[float32] + imaginary[float32] )
    - complex128 ( real[float64] + imaginary[float64] )
- alias
    - byte (alias for unsigned int)
    - rune (alias for unicode code point)

### Zero values
| Data Type | Zero value |
------------ | ------------- |
|int family     | 0 |
|uint family    | 0 |
|complex family | (0+0i) |
|string         | "" (empty string) |
|bool           | false |
|byte           | 0 |
|interface      | nil |
|pointer        | nil |
|function       | nil |
|struct         | struct instance |

## Variables
### Function Scope
- Can use ":="
- CANNOT have unused variables

### Package Scope
- Cannot use ":="
- CAN have unused variables