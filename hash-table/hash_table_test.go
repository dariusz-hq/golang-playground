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

func TestHashTableCollision(t *testing.T) {
	ht := NewHashTable(2)

	someProp1 := "someProp1"
	someProp1Value := 1
	ht.set(someProp1, someProp1Value)

	someProp2 := "someProp2"
	someProp2Value := "some string value"
	ht.set(someProp2, someProp2Value)

	someProp3 := "someProp3"
	someProp3Value := 3.14
	ht.set(someProp3, someProp3Value)

	if ht.get(someProp1) != someProp1Value {
		t.Fatalf("Property %s does not match %d", someProp1, someProp1Value)
	}

	if ht.get(someProp2) != someProp2Value {
		t.Fatalf("Property %s does not match %s", someProp2, someProp2Value)
	}

	if ht.get(someProp3) != someProp3Value {
		t.Fatalf("Property %s does not match %f", someProp3, someProp3Value)
	}
}
