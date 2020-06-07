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

func ReverseTransArray(orgArr []*Transaction) []*Transaction {
	var arr []*Transaction
	if len(orgArr) == 0 {
		return arr
	}
	for i := len(orgArr) - 1; i >= 0; i-- {
		arr = append(arr, orgArr[i])
	}
	return arr
}

func ReverseUTXOArray(orgArr []*UTXO) []*UTXO {
	var arr []*UTXO
	if len(orgArr) == 0 {
		return arr
	}
	for i := len(orgArr) - 1; i >= 0; i-- {
		arr = append(arr, orgArr[i])
	}
	return arr
}
