package utils

import (
	"encoding/json"
	"log"
)

func LogAPICall(apiName string, data interface{}) {
	jsonData, _ := json.Marshal(data)
	log.Printf("API called: %s, Data: %s", apiName, string(jsonData))
}
