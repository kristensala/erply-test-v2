package helpers

import "encoding/json"

func JsonToMap(object any) map[string]string {
	var result map[string]string
    data, _ := json.Marshal(object)
    json.Unmarshal(data, &result)
    return result
}
