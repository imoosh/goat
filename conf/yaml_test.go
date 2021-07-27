package conf

import "testing"

func TestYaml_Parse(t *testing.T) {
    NewYamlParser().Parse("./test/config.yml")
}