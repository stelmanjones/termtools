# Usure

[![Go Reference](https://pkg.go.dev/badge/github.com/username/repo)](https://pkg.go.dev/github.com/stelmanjones/termtools/usure)
[![Go Report Card](https://goreportcard.com/badge/github.com/username/repo)](https://goreportcard.com/report/github.com/stelmanjones/termtools/usure)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/stelmanjones/termtools/blob/main/LICENSE)

A tiny and simple testing library

**Example**
![Example](../examples/images/usure.png)
 
**Output** 
![Output](../examples/images/output.png)

## Install
```go
import "github.com/stelmanjones/termtools/usure"
```

## Features

- **Nil Check**: Checks if the provided value is nil.

```go
result := usure.Nil(value)
```

- **Not Nil Check**: Checks if the provided value is not nil.

```go
result := usure.NotNil(value)
```

- **Instance Check**: Checks if the two provided values are of the same type.

```go
result := usure.IsInstance(value1, value2)
```

- **Equality Check**: Checks if the two provided values are equal.

```go
result := usure.Equal(value1, value2)
```

- **Inequality Check**: Checks if the two provided values are not equal.

```go
result := usure.NotEqual(value1, value2)
```

- **Expect Equality**: Logs an error message if the two provided values are not equal.

```go
usure.ExpectEqual("Values should be equal", value1, value2)
```

