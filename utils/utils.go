package utils

import (
	"encoding/json"
	"flag"
	"fmt"
)

func ToJson(object any) string {
	jsonData, error := json.Marshal(object)
	if error != nil {
		return fmt.Sprintf("error while parse : %s", error.Error())
	}
	return string(jsonData)
}

func LoadEnvFlags() string {
	var env = flag.String("env", "dev", "run your application with env config file!")
	flag.Parse()
	return *env
}
