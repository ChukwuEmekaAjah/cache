package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// KeyValue represents how we would store values internally
type KeyValue struct {
	Key     string   `json:"Key"`
	Value   []string `json:"Value"`
	Command string   `json:"Command"`
}

func parseSetFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {
	keyValueObject := new(KeyValue)
	keyValueObject.Key = arguments[0]
	keyValueObject.Value = arguments[1:]
	keyValueObject.Command = command

	return keyValueObject
}

func parseHSetFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {
	keyValueObject, exists := cacheMap[arguments[0]]

	if !exists {
		keyValueObject := new(KeyValue)
		keyValueObject.Key = arguments[0]
		keyValueObject.Value = arguments[1:]
		keyValueObject.Command = command

		return keyValueObject
	}

	// convert to map
	keyValueMap := make(map[string]string)
	for i := 0; i < len(keyValueObject.Value); i += 2 {
		keyValueMap[keyValueObject.Value[i]] = keyValueObject.Value[i+1]
	}

	// update the map and overwrite existing keys with their respective values
	keyValueMap[arguments[1]] = arguments[2]

	flat := []string{}
	for key, value := range keyValueMap {
		flat = append(flat, key)
		flat = append(flat, value)
	}

	keyValueObject.Value = flat

	return keyValueObject
}

func parseSaddFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if !exists {
		keyValueObject = new(KeyValue)
		keyValueObject.Key = arguments[0]
		keyValueObject.Value = arguments[1:]
		keyValueObject.Command = command

		return keyValueObject
	}

	set := make(map[string]struct{})

	for i := 0; i < len(keyValueObject.Value); i++ {
		set[keyValueObject.Value[i]] = struct{}{}
	}

	for _, arg := range arguments[1:] {
		_, ok := set[arg]
		if !ok {
			set[arg] = struct{}{}
		}
	}

	flat := []string{}
	for key := range set {
		flat = append(flat, key)
	}

	keyValueObject.Value = flat

	return keyValueObject
}

func parseZaddFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if !exists {
		keyValueObject = new(KeyValue)
		keyValueObject.Key = arguments[0]
		keyValueObject.Value = arguments[1:]
		keyValueObject.Command = command

		return keyValueObject
	}

	for i, newValue := range arguments[1:] {
		for u, value := range keyValueObject.Value {
			if i%2 == 0 && u%2 == 0 && value == newValue { // key-value pair already in the hash
				keyValueObject.Value[u+1] = arguments[i+2]
				continue
			}
		}

		keyValueObject.Value = append(keyValueObject.Value, arguments[i+1:i+3]...) // add key-value pair to array
	}

	return keyValueObject
}

func parseLsetFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if !exists {
		keyValueObject = new(KeyValue)
		keyValueObject.Key = arguments[0]
		keyValueObject.Value = arguments[1:]
		keyValueObject.Command = command

		return keyValueObject
	}

	index, err := strconv.ParseInt(arguments[1], 10, 64)

	if err != nil {
		return keyValueObject
	}

	if index > int64(len(keyValueObject.Value)) {
		keyValueObject.Value = append(keyValueObject.Value, arguments[2]) // insert value at position index
	} else {
		keyValueObject.Value[index] = arguments[2]
	}

	return keyValueObject
}

func parseLpushFunction(command string, arguments []string, cacheMap map[string]*KeyValue) *KeyValue {

	keyValueObject, exists := cacheMap[arguments[0]]

	if !exists {
		keyValueObject = new(KeyValue)
		keyValueObject.Key = arguments[0]
		keyValueObject.Value = arguments[1:]
		keyValueObject.Command = command

		return keyValueObject
	}

	keyValueObject.Value = append(keyValueObject.Value, arguments[1:]...)

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

	return strings.Join(keyValue.Value, " "), nil
}

func retrieveHGetFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	keyValue, exists := cacheMap[arguments[0]]

	if !exists {
		return "", errors.New("unable to find value")
	}

	if keyValue.Command != "HSET" {
		return "", errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
	}

	var ans string

	for i, value := range keyValue.Value {
		if value == arguments[1] {
			ans = keyValue.Value[i+1]
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
	noOfFieldsInHash := 0
	keyValue, exists := cacheMap[arguments[0]]

	if !exists {
		return "0", nil
	}

	for i := range keyValue.Value {
		if i%2 == 0 {
			noOfFieldsInHash++
		}
	}

	return strconv.Itoa(noOfFieldsInHash), nil
}

func retrieveLlenFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	value, exists := cacheMap[arguments[0]]

	if !exists {
		return "", errors.New("unable to find key")
	}

	return fmt.Sprint(len(value.Value)), nil
}

func retrieveStrlenFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	value, exists := cacheMap[arguments[0]]

	if !exists {
		return "0", nil
	}

	return fmt.Sprint(len(value.Value[0])), nil
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
	for i, value := range keyValue.Value {
		if i%2 == 0 {
			keys += fmt.Sprintf("%v) %v\t", num, value)
			num++
		}
	}

	return keys, nil
}

func retrievePingFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	var ans string
	var err error

	switch argumentLength := len(arguments); argumentLength {
	case 0:
		ans, err = "PONG", nil
	case 1:
		ans, err = strings.Join(arguments[0:], " "), nil
	default:
		ans, err = "", errors.New("wrong number of arguments for 'ping' command")
	}

	return ans, err
}

func retrieveHExistsFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	keyValue, exists := cacheMap[arguments[0]]

	if !exists {
		return "0", nil
	}

	exists = false
	for i := 0; i < len(keyValue.Value); i += 2 {
		if keyValue.Value[i] == arguments[1] {
			exists = true
			break
		}
	}

	if exists {
		return "1", nil
	} else {
		return "0", nil
	}
}


func retrieveLPopFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	keyValue, exists := cacheMap[arguments[0]]
	var index int64 = 1

	var err error

	if !exists {
		return "(nil)", nil
	}

	if len(arguments) == 2 {
		index, err = strconv.ParseInt(arguments[1], 10, 64)
		if err != nil {
			return "", errors.New("error parsing")
		}
	}

	ans := ""
	num := 1
	for _, v := range keyValue.Value[:index] {
		ans += fmt.Sprintf("%v) %v\t", v, num)
		num++
	}

	keyValue.Value = keyValue.Value[index:]

	return ans, nil
}

func retrieveLindexFunction(commandName string, arguments []string, cacheMap map[string]*KeyValue) (string, error) {
	keyValue, exists := cacheMap[arguments[0]]

	if !exists {
		return "(nil)", nil
	}

	index, err := strconv.ParseInt(arguments[1], 10, 64)
	if err != nil {
		return "", errors.New("error parsing")
	}

	// for -ve indexes
	if index < 0 {
		index = int64(len(keyValue.Value)) - (index * -1)
	}

	return keyValue.Value[index], nil
}

// RetrievalFunctions lists all the functions this cache would support for retrieving values
var RetrievalFunctions = map[string]func(commandKey string, arguments []string, cacheMap map[string]*KeyValue) (string, error){
	"GET":     retrieveGetFunction,
	"HGET":    retrieveHGetFunction,
	"HKEYS":   retrieveHKeysFunction,
	"LPOP":    retrieveLPopFunction,
	"LINDEX":  retrieveLindexFunction,
	"GETSET":  retrieveGetFunction,
	"HGETALL": retrieveGetFunction,
	"HLEN":    retrieveHlenFunction,
	"HMGET":   retrieveGetFunction,
	"PING":    retrievePingFunction,
	"HEXISTS": retrieveHExistsFunction,
	"EXISTS":  retrieveExistsFunction,
	"LLEN":    retrieveLlenFunction,
	"MGET":    retrieveGetFunction,
	"STRLEN":  retrieveStrlenFunction,
	"ZCARD":   retrieveGetFunction,
	"KEYS":    retrieveKeysFunction,
}
