package ipt

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
)

// Option represents a label-value pair for the selection
// Value is always a string, Data is optional extension data
type Option struct {
	Label string
	Value string
	Data  any // Optional extension data, can be nil
}

// Ipt prompts the user with a list of options and returns the value of the selected option.
func Ipt(message string, options []Option, opts ...string) (string, error) {
	if len(options) == 0 {
		return "", errors.New("options cannot be empty")
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
			if opt.Value == defaultValue {
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
		return "", err
	}

	// Find the corresponding value based on the selected label
	for _, opt := range options {
		if opt.Label == selectedLabel {
			return opt.Value, nil
		}
	}

	return "", errors.New("invalid selection")
}

// IptWithData prompts the user with options and returns both value and data
func IptWithData(message string, options []Option, defaultValue string) (string, any, error) {
	if len(options) == 0 {
		return "", nil, errors.New("options cannot be empty")
	}

	// Extract labels for display
	labels := make([]string, len(options))
	for i, opt := range options {
		labels[i] = opt.Label
	}

	// Store the selected label
	var selectedLabel string

	// Find default label if default value is provided
	defaultLabel := labels[0] // Set first option as default if no default value provided
	for _, opt := range options {
		if opt.Value == defaultValue {
			defaultLabel = opt.Label
			break
		}
	}

	// Use survey to ask the question
	err := survey.AskOne(&survey.Select{
		Message: message,
		Options: labels,
		Default: defaultLabel,
	}, &selectedLabel)

	if err != nil {
		return "", nil, err
	}

	// Find the corresponding option based on the selected label
	for _, opt := range options {
		if opt.Label == selectedLabel {
			return opt.Value, opt.Data, nil
		}
	}

	return "", nil, errors.New("invalid selection")
}

func IptRaw(message string, options []string, opts ...string) (string, error) {
	convertedOptions := convertOptions(options)
	return Ipt(message, convertedOptions, opts...)
}

func convertOptions(options []string) []Option {
	var result []Option
	for _, opt := range options {
		result = append(result, Option{Label: opt, Value: opt})
	}
	return result
}
