package BLC

import "encoding/json"

func JsonToArray(data string) ([]string, error) {
	var arr []string
	err := json.Unmarshal([]byte(data), &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}
