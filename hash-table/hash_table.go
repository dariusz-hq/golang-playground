package main

import "fmt"

type HashValue interface {
}

type HashTable struct {
	data map[int]HashValue
	size int
}

func NewHashTable(size int) HashTable {
	var ht HashTable
	ht.size = size
	ht.data = make(map[int]HashValue)
	return ht
}

func (ht HashTable) set(p string, v HashValue) {
	ht.data[ht.hash(p)] = v
}

func (ht HashTable) get(p string) HashValue {
	return ht.data[ht.hash(p)]
}

func (ht HashTable) hash(key string) int {
	hash := 0
	for pos, char := range key {
		hash = (hash + int(char)*pos) % ht.size
	}
	fmt.Println(hash)
	return hash
}
