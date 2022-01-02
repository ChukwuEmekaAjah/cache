package parser

import (
	"testing"
)

func TestWrite(t *testing.T) {
	cache := make(map[string]*KeyValue)

	newValue := KeyValue{
		Key:     "name",
		Value:   []string{"Chukus"},
		Command: "set",
	}

	cache["name"] = &newValue

	_, err := Write(cache)

	if err != nil {
		t.Log("Could not read write file", err)
		t.Fail()
	}

}

func TestRead(t *testing.T) {
	cache, err := Read()

	if err != nil {
		t.Log("Could not read write file", err)
		t.Fail()
	}

	key := "name"
	if _, ok := cache[key]; !ok {
		t.Log("Expected key not in cache map", key)
		t.Fail()
	}
}
