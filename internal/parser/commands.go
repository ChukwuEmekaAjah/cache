package parser

import "strconv"

func isValidSetCommand(command string, arguments []string) bool {
	if len(arguments) < 2 {
		return false
	}

	return true
}

func isValidGetCommand(command string, arguments []string) bool {
	if len(arguments) < 1 {
		return false
	}

	return true
}

func isValidSetexCommand(command string, arguments []string) bool {
	if len(arguments) < 3 {
		return false
	}

	_, err := strconv.ParseInt(arguments[1], 10, 64)

	if err != nil {
		return false
	}

	return true
}

func isValidSaddCommand(command string, arguments []string) bool {
	if len(arguments) < 2 {
		return false
	}

	return true
}

func isValidSetnxCommand(command string, arguments []string) bool {
	if len(arguments) < 2 {
		return false
	}

	return true
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
	if len(arguments) != 3 {
		return false
	}

	return true
}

func isValidLsetCommand(command string, arguments []string) bool {
	if len(arguments) != 3 {
		return false
	}

	_, err := strconv.ParseInt(arguments[1], 10, 64)

	if err != nil {
		return false
	}

	return true
}

func isValidLpushCommand(command string, arguments []string) bool {
	if len(arguments) < 2 {
		return false
	}

	return true
}

func isValidMsetCommand(command string, arguments []string) bool {
	// should have an equal number of keys and values
	if len(arguments)%2 == 1 {
		return false
	}
	return true
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
	if len(arguments) != 2 {
		return false
	}
	return true
}

func isValidHkeysCommand(command string, arguments []string) bool {
	if len(arguments) != 1 {
		return false
	}
	return true
}

func isValidLpopCommand(command string, arguments []string) bool {
	if len(arguments) != 1 {
		return false
	}
	return true
}

func isValidLindexCommand(command string, arguments []string) bool {
	if len(arguments) != 2 {
		return false
	}
	return true
}

func isValidGetSetCommand(command string, arguments []string) bool {
	if len(arguments) != 2 {
		return false
	}
	return true
}

func isValidHgetAllCommand(command string, arguments []string) bool {
	if len(arguments) != 1 {
		return false
	}
	return true
}

func isValidHlenCommand(command string, arguments []string) bool {
	if len(arguments) != 1 {
		return false
	}
	return true
}

func isValidHmgetCommand(command string, arguments []string) bool {
	if len(arguments) < 2 {
		return false
	}
	return true
}

func isValidPingCommand(command string, arguments []string) bool {
	if len(arguments) != 0 {
		return false
	}
	return true
}

func isValidHexistsCommand(command string, arguments []string) bool {
	if len(arguments) != 2 {
		return false
	}
	return true
}

func isValidLlenCommand(command string, arguments []string) bool {
	if len(arguments) != 1 {
		return false
	}
	return true
}

func isValidMgetCommand(command string, arguments []string) bool {
	if len(arguments) < 1 {
		return false
	}
	return true
}

func isValidStrLenCommand(command string, arguments []string) bool {
	if len(arguments) != 1 {
		return false
	}
	return true
}

func isValidZcardCommand(command string, arguments []string) bool {
	if len(arguments) != 1 {
		return false
	}
	return true
}

func isValidExistsCommand(command string, arguments []string) bool {
	if len(arguments) != 1 {
		return false
	}
	return true
}

func isValidKeysCommand(command string, arguments []string) bool {
	return true
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
}
