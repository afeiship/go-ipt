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

	color, err := ipt.Ipt("What is your favorite color?", opts)
	if err != nil {
		f.Error(err)
	}
	f.Log(color)
}

func TestIptRawString(f *testing.T) {
	opts := []ipt.Option[string]{
		{Label: "Red", Value: "Red"},
		{Label: "Green", Value: "Green"},
		{Label: "Blue", Value: "Blue"},
	}

	color, err := ipt.Ipt("What is your favorite color?", opts)
	if err != nil {
		f.Error(err)
	}
	f.Log(color)
}

func TestIptRawMethod(f *testing.T) {
	color, err := ipt.IptRaw("What is your favorite color?", []string{"Red", "Green", "Blue"})
	if err != nil {
		f.Error(err)
	}
	f.Log(color)
}
