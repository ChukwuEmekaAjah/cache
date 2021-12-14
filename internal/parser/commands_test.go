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

func TestGetCommandPresence(t *testing.T) {
	_, commandExists := Commands["GET"]

	if !commandExists {
		t.Log("Command does not exist")
		t.Fail()
	}
}

func TestGetCommandValidity(t *testing.T) {
	command := "get"
	commandArgs := []string{"fine boy"}

	if !isValidGetCommand(command, commandArgs) {
		t.Log("Command should have only an argument")
		t.Fail()
	}
}

func TestExistsCommandPresence(t *testing.T) {
	_, commandExists := Commands["EXISTS"]

	if !commandExists {
		t.Log("Command does not exist")
		t.Fail()
	}
}

func TestExistsValidity(t *testing.T) {
	command := "exists"
	commandArgs := []string{"ajah"}

	if !isValidExistsCommand(command, commandArgs) {
		t.Log("Command should have at least an argument")
		t.Fail()
	}

}