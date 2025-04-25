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
