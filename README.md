<div align="center">
  <br />
  <img width="100" src="__tests__/img.png" alt="go-ipt logo" />
  <br /><br />
  <h1>go-ipt</h1>
  <p><em>Interactive Pipe To: An interactive CLI selection tool for Go</em></p>
</div>

---

## ✨ Features

- Simple CLI prompts with list selection
- Supports default selection
- Extensible `Data` field for attaching metadata
- Cross-platform (uses standard `os` and `bufio`)

## 📦 Installation

```sh
go install github.com/afeiship/go-ipt@latest
```

> 💡 Use `go install` (Go 1.16+) instead of `go get -u` for installing binaries.

## 🚀 Usage

### Basic Example

```go
package main

import "github.com/afeiship/go-ipt"

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
	println(color) // e.g., "red", "green", or "blue"
}
```

### With Default Value

```go
color, err := ipt.Ipt("Select your favorite color:", opts, "green")
// If user presses Enter without selection, returns "green"
```

### Accessing Extended Data

Attach arbitrary metadata using the `Data` field:

```go
type ColorInfo struct {
	Name string
	Hex  string
}

opts := []ipt.Option{
	{Label: "Red", Value: "red", Data: ColorInfo{"Red", "#FF0000"}},
	{Label: "Green", Value: "green", Data: ColorInfo{"Green", "#00FF00"}},
	{Label: "Blue", Value: "blue", Data: ColorInfo{"Blue", "#0000FF"}},
}

value, data, err := ipt.IptWithData("Select color:", opts, "green")
if err != nil {
	panic(err)
}

fmt.Printf("Value: %s\n", value)
if info, ok := data.(ColorInfo); ok {
	fmt.Printf("Name: %s, Hex: %s\n", info.Name, info.Hex)
}
```

### Simplified String List (`IptRaw`)

When label and value are identical:

```go
color, err := ipt.IptRaw("Pick a color:", []string{"Red", "Green", "Blue"}, "Green")
// Default is "Green"
```

## 📚 API Reference

### `func Ipt(message string, options []Option, defaultValue ...string) (string, error)`

Prompts user to select from a list of options.

- `message`: Prompt text
- `options`: Slice of `Option` structs
- `defaultValue...`: Optional default value (uses first element if provided)

Returns selected **value** or error.

---

### `func IptWithData(message string, options []Option, defaultValue string) (string, any, error)`

Same as `Ipt`, but also returns the associated `Data` field.

Returns: `(value, data, error)`

---

### `func IptRaw(message string, options []string, defaultValue ...string) (string, error)`

Convenience function for simple string lists.  
Label and value are the same.

---

### `type Option struct`

```go
type Option struct {
	Label string      // Display label in the prompt
	Value string      // Returned value when selected
	Data  any         // Optional metadata (can be struct, string, etc.)
}
```

## 📝 Notes

- Input is read from `os.Stdin`.
- Uses line-based selection (type number or use arrow keys if supported by terminal).
- If no default is provided and user cancels (e.g., Ctrl+C), returns an error.

## 🛠 License

MIT © [afeiship](https://github.com/afeiship)