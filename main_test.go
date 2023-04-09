package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	length := 10
	str := randomString(length)
	if len(str) != length {
		t.Errorf("Expected string of length %d, got %d", length, len(str))
	}

	if strings.ContainsAny(str, ".,;:/?!") {
		t.Errorf("Generated string contains unexpected characters")
	}
}

func TestUpdateValues(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	jsonData := `{
		"key1": "value1",
		"key2": {
			"nestedKey1": "nestedValue1",
			"nestedKey2": 42,
			"nestedKey3": true,
			"nestedKey4": null
		}
	}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		t.Fatalf("Failed to unmarshal test JSON data: %v", err)
	}

	updateValues(data)

	if key1, ok := data["key1"].(string); !ok || len(key1) != len("value1") {
		t.Errorf("Expected 'key1' to be a random string of length %d, got '%s'", len("value1"), key1)
	}

	key2 := data["key2"].(map[string]interface{})
	if nestedKey1, ok := key2["nestedKey1"].(string); !ok || len(nestedKey1) != len("nestedValue1") {
		t.Errorf("Expected 'nestedKey1' to be a random string of length %d, got '%s'", len("nestedValue1"), nestedKey1)
	}

	if nestedKey2, ok := key2["nestedKey2"].(float64); !ok || nestedKey2 == 42 {
		t.Errorf("Expected 'nestedKey2' to be doubled, got %f", nestedKey2)
	}

	if _, ok := key2["nestedKey3"].(bool); !ok {
		t.Errorf("Expected 'nestedKey3' to be a boolean")
	}

	if key2["nestedKey4"] != nil {
		t.Errorf("Expected 'nestedKey4' to remain null")
	}
}

func TestMainAndOutput(t *testing.T) {
	jsonData := `{
		"key1": "value1",
		"key2": {
			"nestedKey1": "nestedValue1",
			"nestedKey2": 42,
			"nestedKey3": true,
			"nestedKey4": null
		}
	}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		t.Fatalf("Failed to unmarshal test JSON data: %v", err)
	}

	updateValues(data)

	modifiedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal modified JSON data: %v", err)
	}

	var modifiedData map[string]interface{}
	err = json.Unmarshal(modifiedJSON, &modifiedData)
	if err != nil {
		t.Fatalf("Failed to unmarshal modified JSON data: %v", err)
	}
}

func TestRandomValueWithSameDigits(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	cases := []struct {
		name  string
		input float64
	}{
		{"Positive single digit", 7},
		{"Positive two digits", 42},
		{"Positive three digits", 123},
		{"Positive four digits", 1234},
		{"Zero", 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := randomValueWithSameDigits(tc.input)
			if result == tc.input {
				t.Errorf("Expected a different number from input, got the same: %f", result)
			}

			if tc.input != 0 && len(fmt.Sprint(result)) != len(fmt.Sprint(tc.input)) {
				t.Errorf("Expected a number with the same number of digits as input, got: %f", result)
			}
		})
	}
}
