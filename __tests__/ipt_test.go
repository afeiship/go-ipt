package ipt

import (
	"testing"

	"github.com/afeiship/go-ipt"
)

func TestIpt(f *testing.T) {
	// Test with simple string options
	opts := []ipt.Option{
		{Label: "Red", Value: "red"},
		{Label: "Green", Value: "green"},
		{Label: "Blue", Value: "blue"},
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
	colorWithDefault, err := ipt.Ipt("What is your favorite color?", opts, "green")
	if err != nil {
		// Skip test if EOF error (expected during automated testing)
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	f.Log("With default:", colorWithDefault)
}

func TestIptWithData(f *testing.T) {
	type ColorObj struct {
		Name string
		Hex  string
	}

	// Test with data field
	opts := []ipt.Option{
		{Label: "Red", Value: "red", Data: ColorObj{"Red", "#FF0000"}},
		{Label: "Green", Value: "green", Data: ColorObj{"Green", "#00FF00"}},
		{Label: "Blue", Value: "blue", Data: ColorObj{"Blue", "#0000FF"}},
	}

	// Test IptWithDefault function
	value, data, err := ipt.IptWithDefault("What is your favorite color?", opts, "green")
	if err != nil {
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	f.Log("Value:", value, "Data:", data)

	// Verify data is correct
	if colorObj, ok := data.(ColorObj); ok {
		if colorObj.Name != "Green" || colorObj.Hex != "#00FF00" {
			f.Error("Data doesn't match expected values")
		}
	}
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
	opts := []ipt.Option{
		{Label: "Red", Value: "red"},
		{Label: "Green", Value: "green"},
		{Label: "Blue", Value: "blue"},
	}

	// Test with default value not in options - this should not error, just use first option
	color, err := ipt.Ipt("What is your favorite color?", opts, "yellow")
	if err != nil {
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	// Should return first option's value since yellow is not found
	if color != "red" {
		f.Error("Expected first option when default not found")
	}
}

func TestIptWithNilData(f *testing.T) {
	// Test with nil data
	opts := []ipt.Option{
		{Label: "Option1", Value: "opt1", Data: nil},
		{Label: "Option2", Value: "opt2", Data: "simple data"},
		{Label: "Option3", Value: "opt3"}, // Data omitted, should be nil
	}

	// Test Ipt function (ignores data)
	selectedStr, err := ipt.Ipt("Select an option:", opts, "opt2")
	if err != nil {
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}
	if selectedStr != "opt2" {
		f.Error("Expected default value to match")
	}

	// Test IptWithDefault function (includes data)
	value, data, err := ipt.IptWithDefault("Select an option:", opts, "opt2")
	if err != nil {
		if err.Error() == "EOF" {
			f.Skip("Skipping test due to EOF (expected in non-interactive environment)")
		}
		f.Error(err)
	}

	// Verify we got the right value and data
	if value != "opt2" {
		f.Error("Expected opt2 value")
	}
	if data != "simple data" {
		f.Error("Expected 'simple data' for opt2")
	}
}

func TestIptEmptyOptions(f *testing.T) {
	// Test with empty options
	opts := []ipt.Option{}

	_, err := ipt.Ipt("This should fail", opts)
	if err == nil {
		f.Error("Expected error for empty options")
	}

	_, _, err = ipt.IptWithDefault("This should fail", opts, "default")
	if err == nil {
		f.Error("Expected error for empty options in IptWithDefault")
	}
}
