package tool

import (
	"testing"
)

func TestMetaData_Travel(t *testing.T) {
	json, err := FromJsonFile("./1.json")
	//json, err := FromJson([]byte(json1))
	if err != nil {
		t.Fatal(err)
	}
	Travel(json)
}
