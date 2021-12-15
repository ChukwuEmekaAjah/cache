package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// KeyValue represents how we would store values internally
type KeyValue struct {
	key     string
	value   []string
	command string
}

func parseSetFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {
	keyValueObject := new(KeyValue)
	keyValueObject.key = arguments[0]
	keyValueObject.value = arguments[1:]
	keyValueObject.command = command

	return keyValueObject
}

func parseHSetFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {
	keyValueObject, exists := cacheMap[arguments[0]]

	if !exists {
		keyValueObject := new(KeyValue)
		keyValueObject.key = arguments[0]
		keyValueObject.value = arguments[1:]
		keyValueObject.command = command
	
		return keyValueObject
	}

	for _, newValue := range arguments[1:] {
		for _, value := range keyValueObject.value {
			if value == newValue { // value already in the set
				continue
			}
		}
		keyValueObject.value = append(keyValueObject.value, newValue)
	}

	return keyValueObject
}

func parseSaddFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if exists == false {
		keyValueObject = new(KeyValue)
		keyValueObject.key = arguments[0]
		keyValueObject.value = arguments[1:]
		keyValueObject.command = command

		return keyValueObject
	}

	for _, newValue := range arguments[1:] {
		for _, value := range keyValueObject.value {
			if value == newValue { // value already in the set
				continue
			}
		}
		keyValueObject.value = append(keyValueObject.value, newValue)
	}

	return keyValueObject
}

func parseZaddFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if exists == false {
		keyValueObject = new(KeyValue)
		keyValueObject.key = arguments[0]
		keyValueObject.value = arguments[1:]
		keyValueObject.command = command

		return keyValueObject
	}

	for i, newValue := range arguments[1:] {
		for u, value := range keyValueObject.value {
			if i%2 == 0 && u%2 == 0 && value == newValue { // key-value pair already in the hash
				keyValueObject.value[u+1] = arguments[i+2]
				continue
			}
		}

		keyValueObject.value = append(keyValueObject.value, arguments[i+1:i+3]...) // add key-value pair to array
	}

	return keyValueObject
}

func parseLsetFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if exists == false {
		keyValueObject = new(KeyValue)
		keyValueObject.key = arguments[0]
		keyValueObject.value = arguments[1:]
		keyValueObject.command = command

		return keyValueObject
	}

	index, err := strconv.ParseInt(arguments[1], 10, 64)

	if err != nil {
		return keyValueObject
	}

	if index > int64(len(keyValueObject.value)) {
		keyValueObject.value = append(keyValueObject.value, arguments[2]) // insert value at position index
	} else {
		keyValueObject.value[index] = arguments[2]
	}

	return keyValueObject
}

func parseLpushFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if exists == false {
		keyValueObject = new(KeyValue)
		keyValueObject.key = arguments[0]
		keyValueObject.value = arguments[1:]
		keyValueObject.command = command

		return keyValueObject
	}

	keyValueObject.value = append(keyValueObject.value, arguments[1:]...)

	return keyValueObject
}

// ParserFunctions transforms arguments to a Struct
var ParserFunctions = map[string]func(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue{
	"SET":   parseSetFunction,
	"SADD":  parseSaddFunction,
	"SETEX": parseSetFunction,
	"SETNX": parseSetFunction,
	"ZADD":  parseZaddFunction,
	"HSET":  parseHSetFunction,
	"LSET":  parseLsetFunction,
	"LPUSH": parseLpushFunction,
	"MSET":  parseSetFunction,
	"HMSET": parseZaddFunction,
}

func retrieveGetFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	keyValue, exists := cacheMap[arguments[0]]

	if !exists {
		return "", errors.New("unable to find value")
	}

	return strings.Join(keyValue.value, " "), nil
}

func retrieveHGetFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	keyValue, exists := cacheMap[arguments[0]]

	if !exists {
		return "", errors.New("unable to find value")
	}

	var ans string

	for i, value := range keyValue.value {
		if value == arguments[1] {
			ans = keyValue.value[i+1]
		}
	}
	return ans, nil
}

func retrieveExistsFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	noOfKeysPresentCount := 0

	for _, arg := range arguments {
		_, exists := cacheMap[arg]
	
		if exists {
			noOfKeysPresentCount++
		}
	}

	if noOfKeysPresentCount == 0 {
		return "0", errors.New("unable to find key")
	}
	
	return strconv.Itoa(noOfKeysPresentCount), nil
}

func retrieveHlenFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	value, exists := cacheMap[arguments[0]]

	if !exists {
		return "", errors.New("unable to find key")
	}

	return fmt.Sprint(len(value.value) / 2), nil
}

func retrieveLlenFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	value, exists := cacheMap[arguments[0]]

	if exists == false {
		return "", errors.New("Unable to find key")
	}

	return fmt.Sprint(len(value.value)), nil
}

func retrieveStrlenFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	value, exists := cacheMap[arguments[0]]

	if exists == false {
		return "", errors.New("Unable to find key")
	}

	return fmt.Sprint(len(value.value[0])), nil
}

func retrieveKeysFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	keys := ""

	for k := range cacheMap {
		keys += fmt.Sprintf("%s\t", k)
	}

	return keys, nil
}

func retrieveHKeysFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	keys := ""
	keyValue, exists := cacheMap[arguments[0]]

	if !exists {
		return "empty list or set", nil
	}

	num := 1
	for i, value := range keyValue.value {
		if i % 2 == 0 {
			keys += fmt.Sprintf("%v) %v\t", num, value)
			num++
		}
	}

	return keys, nil
}

// RetrievalFunctions lists all the functions this cache would support for retrieving values
var RetrievalFunctions = map[string]func(commandKey string, arguments []string, cacheMap map[string]*KeyValue) (string, error){
	"GET":     retrieveGetFunction,
	"HGET":    retrieveHGetFunction,
	"HKEYS":   retrieveHKeysFunction,
	"LPOP":    retrieveGetFunction,
	"LINDEX":  retrieveGetFunction,
	"GETSET":  retrieveGetFunction,
	"HGETALL": retrieveGetFunction,
	"HLEN":    retrieveHlenFunction,
	"HMGET":   retrieveGetFunction,
	"PING":    retrieveGetFunction,
	"HEXISTS": retrieveGetFunction,
	"EXISTS":  retrieveExistsFunction,
	"LLEN":    retrieveLlenFunction,
	"MGET":    retrieveGetFunction,
	"STRLEN":  retrieveStrlenFunction,
	"ZCARD":   retrieveGetFunction,
	"KEYS":    retrieveKeysFunction,
}
