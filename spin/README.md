# Spin Module

The Spin module provides a thread-safe spinner with various available character sets, prefix, suffix, and color. The spinner can be controlled using an Options pattern.

## Install
```go
import "github.com/stelmanjones/termtools/spin"
```

## Usage

```go
s := spin.New(spin.BouncingBar, time.Millisecond * 10,
    spin.WithPrefix("SPINNING "),
    spin.WithSuffix("AFTER"),
    spin.WithColor(color.FgGreen),
    spin.WithFinalMsg("BYE!"))

s.Start()
time.Sleep(time.Second * 3)
s.Stop()
```

