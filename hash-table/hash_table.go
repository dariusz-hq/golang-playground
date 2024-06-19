package main

import "fmt"

type HashValue interface {
}

type HashTable struct {
	data map[int]map[string]HashValue
	size int
}

func NewHashTable(size int) HashTable {
	var ht HashTable
	ht.size = size
	ht.data = map[int]map[string]HashValue{}
	return ht
}

func (ht HashTable) set(p string, v HashValue) {
	hash := ht.hash(p)
	if ht.data[hash] == nil {
		myMap := map[string]HashValue{p: v}
		ht.data[hash] = myMap
	} else {
		currentMap := ht.data[hash]
		currentMap[p] = v
		ht.data[hash] = currentMap
	}

	fmt.Println(ht)
}

func (ht HashTable) get(p string) HashValue {
	var result HashValue
	hash := ht.hash(p)
	for index, value := range ht.data[hash] {
		if index == p {
			result = value
		}
	}
	return result
}

func (ht HashTable) hash(key string) int {
	hash := 0
	for pos, char := range key {
		hash = (hash + int(char)*pos) % ht.size
	}
	fmt.Println(hash)
	return hash
}
