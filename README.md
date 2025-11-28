<div align="center">
	<br>
	<img width="100" src="__tests__/img.png" alt="ipt logo">
	<br>
</div>

# go-ipt
> Interactive Pipe To: The golang cli interactive workflow.


## installation
```sh
go get -u github.com/afeiship/go-ipt
```

## usage

### Basic usage

```go
package main

import (
	"github.com/afeiship/go-ipt"
)

func main() {
	opts := []ipt.Option{
		{Label: "Red", Value: "red"},
		{Label: "Green", Value: "green"},
		{Label: "Blue", Value: "blue"},
	}

	color, err := ipt.Ipt("What is your favorite color?", opts)
	if err != nil {
		panic(err)
	}
	println(color) // Returns "red", "green", or "blue"
}
```

### With default value

```go
package main

import (
	"github.com/afeiship/go-ipt"
)

func main() {
	// Options with default value
	opts := []ipt.Option{
		{Label: "Red", Value: "red"},
		{Label: "Green", Value: "green"},
		{Label: "Blue", Value: "blue"},
	}

	// Set "Green" as default
	selectedColor, err := ipt.Ipt("Select your favorite color:", opts, "green")
	if err != nil {
		panic(err)
	}
	println(selectedColor) // Will be "green" if user just presses Enter
}
```

### Using Data field for extended information

```go
package main

import (
	"github.com/afeiship/go-ipt"
	"fmt"
)

func main() {
	type ColorInfo struct {
		Name string
		Hex  string
	}

	opts := []ipt.Option{
		{Label: "Red", Value: "red", Data: ColorInfo{"Red", "#FF0000"}},
		{Label: "Green", Value: "green", Data: ColorInfo{"Green", "#00FF00"}},
		{Label: "Blue", Value: "blue", Data: ColorInfo{"Blue", "#0000FF"}},
	}

	// Get both value and data
	value, data, err := ipt.IptWithDefault("Select color:", opts, "green")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value: %s\n", value)
	if colorInfo, ok := data.(ColorInfo); ok {
		fmt.Printf("Name: %s, Hex: %s\n", colorInfo.Name, colorInfo.Hex)
	}
}
```

### Using IptRaw

```go
package main

import (
	"github.com/afeiship/go-ipt"
)

func main() {
	// Simple string options with default
	color, err := ipt.IptRaw("Select color:", []string{"Red", "Green", "Blue"}, "Green")
	if err != nil {
		panic(err)
	}
	println(color) // Will be "Green" if user just presses Enter
}
```

## API

### `Ipt(message string, options []Option, opts ...string) (string, error)`

- `message`: The prompt message to display
- `options`: Array of Option containing Label, Value, and optional Data
- `opts`: Optional default value (first argument if provided)

### `IptWithDefault(message string, options []Option, defaultValue string) (string, any, error)`

- `message`: The prompt message to display
- `options`: Array of Option containing Label, Value, and optional Data
- `defaultValue`: Default value to select
- Returns: selected value, data, and error

### `IptRaw(message string, options []string, opts ...string) (string, error)`

- `message`: The prompt message to display
- `options`: Array of string options (label and value are the same)
- `opts`: Optional default value (first argument if provided)

### `Option struct`

- `Label`: The display text for the option
- `Value`: The string value to return when selected
- `Data`: Optional extension data of any type, can be nil