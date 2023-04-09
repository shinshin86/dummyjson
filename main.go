package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
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
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())

	updateValues(data)

	modifiedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(modifiedJSON))
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func randomValueWithSameDigits(n float64) float64 {
	min := int(math.Pow(10, float64(len(fmt.Sprint(n))-1)))
	max := int(math.Pow(10, float64(len(fmt.Sprint(n)))) - 1)

	for {
		randNum := rand.Intn(max-min+1) + min
		if float64(randNum) != n {
			return float64(randNum)
		}
	}
}

func updateValues(data interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		for k, val := range v {
			switch val := val.(type) {
			case string:
				v[k] = randomString(len(val))
			case float64:
				v[k] = randomValueWithSameDigits(val)
			case bool:
				v[k] = rand.Float32() < 0.5
			case nil:
				// keep null
			default:
				updateValues(val)
			}
		}
	case []interface{}:
		for i, val := range v {
			switch val := val.(type) {
			case string:
				v[i] = randomString(len(val))
			case float64:
				v[i] = val * 2
			case bool:
				v[i] = rand.Float32() < 0.5
			case nil:
				// keep null
			default:
				updateValues(val)
			}
		}
	}
}
