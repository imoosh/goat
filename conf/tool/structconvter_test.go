package tool

import (
	"fmt"
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

func TestMetaData_Travel1(t *testing.T) {
	xml, err := FromXmlFile("./1.xml")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(xml)
	//Travel(xml)
}
