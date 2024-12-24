package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Define the Animal interface
type Animal interface {
}

// Define concrete types: Cat and Dog
type Cat struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Meows bool   `json:"meows"`
}

type Dog struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Barks bool   `json:"barks"`
}

// A helper type for polymorphic deserialization
type AnimalAdapter struct {
	Animal
}

// Implement the custom UnmarshalJSON method
func (a *AnimalAdapter) UnmarshalJSON(data []byte) error {
	// Check the "type" field first
	var wrapper struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return err
	}

	// Based on the type, unmarshal into the correct struct
	switch wrapper.Type {
	case "cat":
		var cat Cat
		if err := json.Unmarshal(data, &cat); err != nil {
			return err
		}
		a.Animal = cat
	case "dog":
		var dog Dog
		if err := json.Unmarshal(data, &dog); err != nil {
			return err
		}
		a.Animal = dog
	default:
		return fmt.Errorf("unknown type: %s", wrapper.Type)
	}

	return nil
}

// Main function
func TestFoo(t *testing.T) {
	jsonData := `{ "type": "cat", "name": "Whiskers", "meows": true }`

	var animalAdapter AnimalAdapter
	if err := json.Unmarshal([]byte(jsonData), &animalAdapter); err != nil {
		t.Fatal("Error:", err)
		return
	}

	switch a := animalAdapter.Animal.(type) {
	case Cat:
		fmt.Printf("meows %v", a.Meows)
	case Dog:
		fmt.Printf("barsk %v", a.Barks)
	default:
		t.Fatalf("Unknown type %s", a)
	}

	equals := cmp.Equal(animalAdapter.Animal, Cat{Type: "cat", Name: "Whiskers", Meows: true})
	if !equals {
		t.Fatalf("Not equal")
	}
}
