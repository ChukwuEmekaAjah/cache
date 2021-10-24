package parser

import (
	"testing"
)

func TestCommands(t *testing.T) {
	_, commandExists := Commands["SET"]

	if !commandExists {
		t.Log("Command does not exist")
		t.Fail()
	}
}

func TestCommandAbsence(t *testing.T) {
	_, commandExists := Commands["HGETTER"]

	if commandExists {
		t.Log("Command should not exist")
		t.Fail()
	}
}

func TestSetexCommandValidity(t *testing.T) {
	command := "setex"
	commandArgs := []string{"name", "32", "ajah"}

	if !isValidSetexCommand(command, commandArgs) {
		t.Log("Command should be a valid SETEX command")
		t.Fail()
	}
}

func TestSetexCommandInvalidity(t *testing.T) {
	command := "setex"
	commandArgs := []string{"name"}

	if isValidSetexCommand(command, commandArgs) {
		t.Log("Command should have at least 2 arguments")
		t.Fail()
	}
}

func TestSetexCommandInvalidExpiryTime(t *testing.T) {
	command := "setex"
	commandArgs := []string{"name", "ajah", "emeka"}

	if isValidSetexCommand(command, commandArgs) {
		t.Log("Command expiry time should be a valid number")
		t.Fail()
	}
}
