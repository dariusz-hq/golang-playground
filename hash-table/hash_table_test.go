package main

import (
	"testing"
)

func TestHashTableSetInt(t *testing.T) {
	ht := NewHashTable(50)
	someIntValue := 45
	somePropName := "someProp"
	ht.set(somePropName, someIntValue)
	if ht.get(somePropName) != someIntValue {
		t.Fatalf("Property %s does not match %d", somePropName, someIntValue)
	}
}

func TestHashTableSetFloat(t *testing.T) {
	ht := NewHashTable(50)
	someFloatValue := 45.1235
	somePropName := "someProp"
	ht.set(somePropName, someFloatValue)
	if ht.get(somePropName) != someFloatValue {
		t.Fatalf("Property %s does not match %f", somePropName, someFloatValue)
	}
}

func TestHashTableSetString(t *testing.T) {
	ht := NewHashTable(50)
	someStringValue := "some test string"
	somePropName := "someProp"
	ht.set(somePropName, someStringValue)
	if ht.get(somePropName) != someStringValue {
		t.Fatalf("Property %s does not match %s", somePropName, someStringValue)
	}
}
