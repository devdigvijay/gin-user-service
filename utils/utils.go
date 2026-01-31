package utils

import (
	"encoding/json"
	"fmt"
)

func ToJson(object any) string {
	jsonData, error := json.Marshal(object)
	if error != nil {
		return fmt.Sprintf("error while parse : %s", error.Error())
	}
	return string(jsonData)
}
