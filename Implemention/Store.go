package Implemention

import "fmt"

var store = make(map[string]string)

func StorePut(key, value string) {
	fmt.Printf("%s: %s", key, value)
	store[key] = value
}

func StoreGet(key string) string {
	return store[key]
}
