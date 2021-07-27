package tool

import "testing"

func TestFromJsonFile(t *testing.T) {
    t.Log(FromJsonFile("./1.json"))
}
