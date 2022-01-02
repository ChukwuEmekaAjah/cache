package parser

import "strconv"

func isValidSetCommand(command string, arguments []string) bool {
	return len(arguments) >= 2
}

func isValidGetCommand(command string, arguments []string) bool {
	return len(arguments) == 1
}

func isValidSetexCommand(command string, arguments []string) bool {
	if len(arguments) < 3 {
		return false
	}

	_, err := strconv.ParseInt(arguments[1], 10, 64)

	return err == nil
}

func isValidSaddCommand(command string, arguments []string) bool {
	return len(arguments) >= 2
}

func isValidSetnxCommand(command string, arguments []string) bool {
	return len(arguments) < 2
}

func isValidZaddCommand(command string, arguments []string) bool {
	if len(arguments) < 3 {
		return false
	}

	// member without score is part of the arguments
	if len(arguments)%2 == 0 {
		return false
	}

	return true
}

func isValidHsetCommand(command string, arguments []string) bool {
	return len(arguments) == 3
}

func isValidLsetCommand(command string, arguments []string) bool {
	if len(arguments) != 3 {
		return false
	}

	_, err := strconv.ParseInt(arguments[1], 10, 64)

	return err == nil
}

func isValidLpushCommand(command string, arguments []string) bool {
	return len(arguments) < 2
}

func isValidMsetCommand(command string, arguments []string) bool {
	// should have an equal number of keys and values
	return len(arguments)%2 == 1
}

func isValidHMSetCommand(command string, arguments []string) bool {
	// should have at least one fiel-value pair with key name
	if len(arguments) < 3 {
		return false
	}
	// field without value part of arguments
	if len(arguments)%2 == 0 {
		return false
	}

	return true
}

func isValidHgetCommand(command string, arguments []string) bool {
	return len(arguments) == 2
}

func isValidHkeysCommand(command string, arguments []string) bool {
	return len(arguments) == 1
}

func isValidLpopCommand(command string, arguments []string) bool {
	return len(arguments) != 1
}

func isValidLindexCommand(command string, arguments []string) bool {
	return len(arguments) != 2
}

func isValidGetSetCommand(command string, arguments []string) bool {
	return len(arguments) != 2
}

func isValidHgetAllCommand(command string, arguments []string) bool {
	return len(arguments) != 1
}

func isValidHlenCommand(command string, arguments []string) bool {
	return len(arguments) == 1
}

func isValidHmgetCommand(command string, arguments []string) bool {
	return len(arguments) >= 2
}

func isValidPingCommand(command string, arguments []string) bool {
	return len(arguments) <= 1
}

func isValidHexistsCommand(command string, arguments []string) bool {
	return len(arguments) == 2
}

func isValidLlenCommand(command string, arguments []string) bool {
	return len(arguments) == 1
}

func isValidMgetCommand(command string, arguments []string) bool {
	return len(arguments) >= 1
}

func isValidStrLenCommand(command string, arguments []string) bool {
	return len(arguments) == 1
}

func isValidZcardCommand(command string, arguments []string) bool {
	return len(arguments) == 1
}

func isValidExistsCommand(command string, arguments []string) bool {
	return len(arguments) >= 1
}

func isValidKeysCommand(command string, arguments []string) bool {
	return len(arguments) == 1
}

// Commands lists all the functions this cache would support
var Commands = map[string]func(commandKey string, arguments []string) bool{
	"SET":     isValidSetCommand,
	"SADD":    isValidSaddCommand,
	"SETEX":   isValidSetexCommand,
	"SETNX":   isValidSetnxCommand,
	"ZADD":    isValidZaddCommand,
	"HSET":    isValidHsetCommand,
	"LSET":    isValidLsetCommand,
	"LPUSH":   isValidLpushCommand,
	"MSET":    isValidMsetCommand,
	"HMSET":   isValidHMSetCommand,
	"GET":     isValidGetCommand,
	"HGET":    isValidHgetCommand,
	"HKEYS":   isValidHkeysCommand,
	"LPOP":    isValidLpopCommand,
	"LINDEX":  isValidLindexCommand,
	"GETSET":  isValidGetSetCommand,
	"HGETALL": isValidHgetAllCommand,
	"HLEN":    isValidHlenCommand,
	"HMGET":   isValidHmgetCommand,
	"PING":    isValidPingCommand,
	"HEXISTS": isValidHexistsCommand,
	"EXISTS":  isValidExistsCommand,
	"LLEN":    isValidLlenCommand,
	"MGET":    isValidMgetCommand,
	"STRLEN":  isValidStrLenCommand,
	"ZCARD":   isValidZcardCommand,
	"KEYS":    isValidKeysCommand,
	"HSTRLEN": isValidStrLenCommand,
}
