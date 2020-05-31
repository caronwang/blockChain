package BLC

import (
	"fmt"
	"testing"
)

/*
	json转数组
*/
func TestJsonToArray(t *testing.T) {
	data := `["nihao","world"]`
	arr, _ := JsonToArray(data)
	for _, v := range arr {
		fmt.Println(v)
	}
}
