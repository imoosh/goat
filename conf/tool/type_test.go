package tool

import (
    "fmt"
    "testing"
)

func TestMetaData_Travel(t *testing.T) {
    json, err := FromJsonFile("./1.json")
    if err != nil {
        t.Fatal(err)
    }
    fmt.Println(json)
    json.Travel()
}
