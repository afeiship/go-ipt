package ipt

import (
	"testing"

	"github.com/afeiship/go-ipt"
)

func TestIpt(f *testing.T) {
	type ColorObj struct {
		Name string
		Hex  string
	}

	opts := []ipt.Option[ColorObj]{
		{Label: "Red", Value: ColorObj{"Red", "#FF0000"}},
		{Label: "Green", Value: ColorObj{"Green", "#00FF00"}},
		{Label: "Blue", Value: ColorObj{"Blue", "#0000FF"}},
	}

	// Test without default value
	color, err := ipt.Ipt("What is your favorite color?", opts)
	if err != nil {
		// Skip test if EOF error (expected during automated testing)
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	f.Log("Without default:", color)

	// Test with default value
	defaultColor := ColorObj{"Green", "#00FF00"}
	colorWithDefault, err := ipt.Ipt("What is your favorite color?", opts, defaultColor)
	if err != nil {
		// Skip test if EOF error (expected during automated testing)
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	f.Log("With default:", colorWithDefault)
}

func TestIptRawString(f *testing.T) {
	opts := []ipt.Option[string]{
		{Label: "Red", Value: "Red"},
		{Label: "Green", Value: "Green"},
		{Label: "Blue", Value: "Blue"},
	}

	// Test without default value
	color, err := ipt.Ipt("What is your favorite color?", opts)
	if err != nil {
		// Skip test if EOF error (expected during automated testing)
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	f.Log("Without default:", color)

	// Test with default value
	colorWithDefault, err := ipt.Ipt("What is your favorite color?", opts, "Blue")
	if err != nil {
		// Skip test if EOF error (expected during automated testing)
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	f.Log("With default:", colorWithDefault)
}

func TestIptRawMethod(f *testing.T) {
	// Test without default value
	color, err := ipt.IptRaw("What is your favorite color?", []string{"Red", "Green", "Blue"})
	if err != nil {
		// Skip test if EOF error (expected during automated testing)
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	f.Log("Without default:", color)

	// Test with default value
	colorWithDefault, err := ipt.IptRaw("What is your favorite color?", []string{"Red", "Green", "Blue"}, "Green")
	if err != nil {
		// Skip test if EOF error (expected during automated testing)
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	f.Log("With default:", colorWithDefault)
}

func TestIptWithInvalidDefault(f *testing.T) {
	opts := []ipt.Option[string]{
		{Label: "Red", Value: "Red"},
		{Label: "Green", Value: "Green"},
		{Label: "Blue", Value: "Blue"},
	}

	// Test with default value not in options
	_, err := ipt.Ipt("What is your favorite color?", opts, "Yellow")
	if err == nil {
		f.Error("Expected error when default value is not in options")
	}
}

func TestIptStructDefault(f *testing.T) {
	type ColorObj struct {
		Name string
		Hex  string
	}

	defaultColor := ColorObj{"Cyan", "#00FFFF"}
	opts := []ipt.Option[ColorObj]{
		{Label: "Red", Value: ColorObj{"Red", "#FF0000"}},
		{Label: "Green", Value: ColorObj{"Green", "#00FF00"}},
		{Label: "Blue", Value: ColorObj{"Blue", "#0000FF"}},
		{Label: "Cyan", Value: defaultColor},
	}

	// Test struct default value
	color, err := ipt.Ipt("What is your favorite color?", opts, defaultColor)
	if err != nil {
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	if color != defaultColor {
		f.Error("Expected default struct value to match")
	}
}

func TestIptSingleSelectWithDefault(f *testing.T) {
	// Test with string default value
	strOpts := []ipt.Option[string]{
		{Label: "Option1", Value: "Value1"},
		{Label: "Option2", Value: "Value2"},
		{Label: "Option3", Value: "Value3"},
	}

	// Test string default
	selectedStr, err := ipt.Ipt("Select an option (string):", strOpts, "Value2")
	if err != nil {
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	if selectedStr != "Value2" {
		f.Error("Expected default string value to match")
	}

	// Test with struct default value
	type TestStruct struct {
		ID   int
		Name string
	}

	defaultStruct := TestStruct{ID: 2, Name: "Default"}
	structOpts := []ipt.Option[TestStruct]{
		{Label: "Struct1", Value: TestStruct{ID: 1, Name: "One"}},
		{Label: "Struct2", Value: defaultStruct},
		{Label: "Struct3", Value: TestStruct{ID: 3, Name: "Three"}},
	}

	// Test struct default
	selectedStruct, err := ipt.Ipt("Select an option (struct):", structOpts, defaultStruct)
	if err != nil {
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	if selectedStruct != defaultStruct {
		f.Error("Expected default struct value to match")
	}
}
