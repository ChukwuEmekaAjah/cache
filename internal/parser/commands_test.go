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

func TestHSetCommand(t *testing.T) {
	_, commandExists := Commands["HSET"]

	if !commandExists {
		t.Log("Command does not exist")
		t.Fail()
	}

	command := "hset"
	commandArgs := []string{"mentor", "name", "ajah"}

	if !isValidHsetCommand(command, commandArgs) {
		t.Log("Wrong number of arguments for 'hset' command")
		t.Fail()
	}
}

func TestHGetCommand(t *testing.T) {
	_, commandExists := Commands["HGET"]

	if !commandExists {
		t.Log("Command does not exist")
		t.Fail()
	}

	command := "hget"
	commandArgs := []string{"mentor", "name"}

	if !isValidHgetCommand(command, commandArgs) {
		t.Log("Wrong number of arguments for 'hget' command")
		t.Fail()
	}
}

func TestHKeysCommand(t *testing.T) {
	command := "hkeys"
	commandArgs := []string{"mentor"}

	if !isValidHkeysCommand(command, commandArgs) {
		t.Log("Wrong number of arguments for 'hkeys' command")
		t.Fail()
	}
}

func TestKeysCommand(t *testing.T) {
	_, commandExists := Commands["KEYS"]

	if !commandExists {
		t.Log("Command does not exist")
		t.Fail()
	}
}

func TestHLenCommands(t *testing.T) {
	_, commandExists := Commands["HLEN"]

	if !commandExists {
		t.Log("Command does not exist")
		t.Fail()
	}

	command := "hlen"
	commandArgs := []string{"mentor"}

	if !isValidHlenCommand(command, commandArgs) {
		t.Log("Wrong number of arguments for 'hlen' command")
		t.Fail()
	}
}

func TestSAddCommand(t *testing.T) {
	_, commandExists := Commands["SADD"]

	if !commandExists {
		t.Log("Command does not exist")
		t.Fail()
	}

	command, commandArgs := "sadd", []string{"myset", "hello1"}
	
	if !isValidSaddCommand(command, commandArgs) {
		t.Log("Wrong number of arguments for 'sadd' command")
		t.Fail()
	}
}

func TestStrLenCommand(t *testing.T) {
	_, commandExists := Commands["STRLEN"]

	if !commandExists {
		t.Log("Command does not exist")
		t.Fail()
	}

	command, commandArgs := "strlen", []string{"name"}
	
	if !isValidStrLenCommand(command, commandArgs) {
		t.Log("Wrong number of arguments for 'strlen' command")
		t.Fail()
	}
}

func TestHExistsCommand(t *testing.T) {
	_, commandExists := Commands["HEXISTS"]

	if !commandExists {
		t.Log("Command does not exist")
		t.Fail()
	}
}
