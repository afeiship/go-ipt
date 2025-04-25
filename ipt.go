package ipt

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
)

// Option represents a label-value pair for the selection
type Option[T any] struct {
	Label string
	Value T
}

// Ipt prompts the user with a list of options and returns the value of the selected option.
func Ipt[T any](message string, options []Option[T], opts ...T) (T, error) {
	var zeroValue T // Default value of type T

	if len(options) == 0 {
		return zeroValue, errors.New("options cannot be empty")
	}

	// Extract labels for display
	labels := make([]string, len(options))
	for i, opt := range options {
		labels[i] = opt.Label
	}

	// Store the selected label
	var selectedLabel string

	// Find default label if default value is provided in opts
	defaultLabel := labels[0] // Set first option as default if no default value provided
	if len(opts) > 0 {
		defaultValue := opts[0]
		for _, opt := range options {
			if interface{}(opt.Value) == interface{}(defaultValue) {
				defaultLabel = opt.Label
				break
			}
		}
	}

	// Use survey to ask the question
	err := survey.AskOne(&survey.Select{
		Message: message,
		Options: labels,
		Default: defaultLabel,
	}, &selectedLabel)

	if err != nil {
		return zeroValue, err
	}

	// Find the corresponding value based on the selected label
	for _, opt := range options {
		if opt.Label == selectedLabel {
			return opt.Value, nil
		}
	}

	return zeroValue, errors.New("invalid selection")
}

func IptRaw(message string, options []string, opts ...string) (string, error) {
	return Ipt[string](message, convertOptions[string](options), opts...)
}

func convertOptions[T any](options []string) []Option[string] {
	var result []Option[string]
	for _, opt := range options {
		result = append(result, Option[string]{Label: opt, Value: opt})
	}
	return result
}
