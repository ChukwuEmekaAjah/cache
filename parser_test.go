package cache

import (
	"testing"
)

func TestParse(t *testing.T) {
	command := "SET name ajah"
	if !Parse(command) {
		t.Log("Failed parsing")
		t.Fail()
	}
}

func TestSetParseWithoutKeyValue(t *testing.T) {
	command := "set name"

	if Parse(command) {
		t.Log("Parse should return false for incomplete arguments")
		t.Fail()
	}
}

func TestInvalidParseCommand(t *testing.T) {
	command := "seter name ajah"
	if Parse(command) {
		t.Log("Failed parsing")
		t.Fail()
	}
}

func TestGetParse(t *testing.T) {
	command := "get name"

	if !Parse(command) {
		t.Log("Failed parsing for GET command")
		t.Fail()
	}
}

func TestGetParseWithoutKey(t *testing.T) {
	command := "get"

	if Parse(command) {
		t.Log("Incorrect parsing of GET command")
		t.Fail()
	}
}
