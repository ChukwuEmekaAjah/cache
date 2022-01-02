package parser

import "strings"

//Parse is the top-level function that accepts client commands on the query
func Parse(command string) bool {
	commandParts := strings.Fields(command)
	commandKey := commandParts[0]
	commandParsingFunction, commandExists := Commands[strings.ToUpper(commandKey)]

	if !commandExists {
		return false
	}

	isValidRequest := commandParsingFunction(commandKey, commandParts[1:])

	return isValidRequest
}
