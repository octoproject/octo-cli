package prompt

import (
	"errors"
	"strconv"
	"strings"

	survey "github.com/AlecAivazis/survey/v2"
)

func PromptConfirm(msg string) (bool, error) {
	var data bool
	prompt := &survey.Confirm{
		Message: msg,
	}
	survey.AskOne(prompt, &data)
	return data, nil
}

func PromptSelect(msg string, opts []string) (string, error) {
	var data string
	prompt := &survey.Select{
		Message: msg,
		Options: opts,
	}
	survey.AskOne(prompt, &data)
	return data, nil
}

func PromptString(msg string, required bool) (string, error) {
	var data string
	prompt := &survey.Input{
		Message: msg,
	}

	err := survey.AskOne(prompt, &data)
	if err != nil {
		return "", err
	}

	if required {
		err := ValidateEmptyInput(data)
		if err != nil {
			return "", err
		}
	}
	return strings.TrimSpace(data), nil
}

func PromptInteger(msg string, required bool) (int, error) {
	var data string
	prompt := &survey.Input{
		Message: msg,
	}

	err := survey.AskOne(prompt, &data)
	if err != nil {
		return 0, err
	}

	if required {
		err := ValidateIntegerNumberInput(data)
		if err != nil {
			return 0, err
		}
	}

	return strconv.Atoi(data)
}

func ValidateEmptyInput(input string) error {
	if len(strings.TrimSpace(input)) < 1 {
		return errors.New("this input must not be empty")
	}
	return nil
}

func ValidateIntegerNumberInput(input string) error {
	_, err := strconv.ParseInt(input, 0, 64)
	if err != nil {
		return errors.New("invalid number")
	}
	return nil
}
