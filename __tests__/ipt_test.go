package ipt

import (
	"testing"

	"github.com/afeiship/go-ipt"
)

func TestSayHi(f *testing.T) {

	type ColorObj struct {
		Name string
		Hex  string
	}

	opts = []ipt.Option{
		{"Red", ColorObj{"红色", "#FF0000"}},
		{"Green", ColorObj{"绿色", "#00FF00"}},
		{"Blue", ColorObj{"蓝色", "#0000FF"}},
	}

	ipt.Ipt("Choose a color: ", opts)
}
