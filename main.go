package main

import (
	"encoding/json"
	"fmt"

	"github.com/webtoor/signature/algorithm"
)

const (
	apiKey    = `IhZFxWTK7cFYaWzSRyz4if3c4L3HTwDjb9lvQajZ`
	secretKey = `4EwABfExAeBgKa212rdUgYB5qNd5vVlHKAocUW9R`
)

type Client struct {
	Name   string `json:"name"`
	ApiKey string `json:"api_key"`
	Status string `json:"status"`
}

func CreateSign(algoType, data string) string {
	switch algoType {
	case "HS256":
		return algorithm.HS256Compute(data, secretKey)
	default:
		return "unsupported algorithm"
	}
}

func main() {
	// initial data
	client := Client{
		Name:   "moonshi",
		ApiKey: apiKey,
		Status: "active",
	}
	jsonByte, _ := json.Marshal(client)
	data := string(jsonByte)
	fmt.Println(data)

	// create signature
	cs := CreateSign("HS256", data)
	fmt.Println(cs)

	// validate authenticity of the data
	v := algorithm.HS256Verify(data, cs, secretKey)
	fmt.Println(v)
}
