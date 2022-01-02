package parser

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//Read is the top-level function that accepts client commands on the query
func Read() (map[string]*KeyValue, error) {

	var cacheMap map[string]*KeyValue
	currentDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	jsonString, err := ioutil.ReadFile(currentDir + string(os.PathSeparator) + "data.json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonString, &cacheMap)

	if err != nil {
		return nil, err
	}

	return cacheMap, nil
}

//Write is the top-level function that accepts client commands on the query
func Write(cacheMap map[string]*KeyValue) (bool, error) {
	jsonString, err := json.Marshal(cacheMap)

	if err != nil {
		return false, err
	}

	currentDir, err := os.Getwd()

	if err != nil {
		return false, err
	}

	err = ioutil.WriteFile(currentDir+string(os.PathSeparator)+"data.json", []byte(jsonString), os.FileMode(os.O_RDWR))

	if err != nil {
		return false, err
	}

	return true, nil
}
