package cache

import (
	"fmt"
	"strings"
	"testing"
)

var cacheMap = make(map[string]*KeyValue)

func TestSetCommandParser(t *testing.T) {
	command := "set name ajah"
	commandParts := strings.Fields(command)

	parsedValue := ParserFunctions[strings.ToUpper(commandParts[0])](strings.ToUpper(commandParts[0]), commandParts[1:], cacheMap)

	cacheMap[commandParts[1]] = parsedValue

	if parsedValue == nil {
		t.Log("Invalid arguments parsed into command")
		t.Fail()
	}
}

func TestGetCommandParser(t *testing.T) {
	command := "get name"
	commandParts := strings.Fields(command)

	parsedValue, err := RetrievalFunctions[strings.ToUpper(commandParts[0])](commandParts[0], commandParts[1:], cacheMap)

	fmt.Printf("value is %v\n", parsedValue)

	if err != nil {
		t.Log("Invalid arguments parsed into command")
		t.Fail()
	}
}

func TestGetCommandParserFailure(t *testing.T) {
	command := "get age"
	commandParts := strings.Fields(command)

	_, err := RetrievalFunctions[strings.ToUpper(commandParts[0])](commandParts[0], commandParts[1:], cacheMap)

	if err == nil {
		t.Log("Command should not return a value for a non-existent key")
		t.Fail()
	}
}

func TestKeysCommandParser(t *testing.T) {
	command := []string{"set person ajah", "set name chuks", "set age 32"}

	for _, v := range command {
		commandParts := strings.Fields(v)
		parsedValue := ParserFunctions[strings.ToUpper(commandParts[0])](strings.ToUpper(commandParts[0]), commandParts[1:], cacheMap)
		cacheMap[commandParts[1]] = parsedValue
	}

	comm := "keys"
	commandParts := strings.Fields(comm)

	d, err := RetrievalFunctions[strings.ToUpper(commandParts[0])](commandParts[0], commandParts[1:], cacheMap)

	if err != nil {
		t.Log("Command should return set keys")
		t.Fail()
	}

	for k := range cacheMap {

		if !strings.Contains(d, k) {
			t.Log("Incorrect keys returned")
			t.Fail()
		}
	}
}
